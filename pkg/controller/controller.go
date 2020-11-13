package controller

import (
	"boiler/pkg/logging"
	"boiler/pkg/requests"
	"boiler/pkg/source"
	"boiler/pkg/transformation"
	"context"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/errgroup"
	"time"
)

type BudgetConfig struct {
	TimeBudget time.Duration
}

type Config struct {
	Concurrency     int
	ContinueOnError bool
	Budget          BudgetConfig
}

type Controller struct {
	source          source.Source
	transformations []transformation.Transformation
	executor        Executor
	config          Config
	logger          logging.Logger
}

func NewController(source source.Source, transformations []transformation.Transformation, executor Executor, config Config, logger logging.Logger) Controller {
	return Controller{
		source:          source,
		transformations: transformations,
		executor:        executor,
		config:          config,
		logger:          logger,
	}
}

func (c Controller) Execute(parentCtx context.Context) error {
	reqsChan := make(chan requests.Request, c.config.Concurrency)

	var cancel func()
	if c.config.Budget.TimeBudget != 0 {
		parentCtx, cancel = context.WithTimeout(parentCtx, c.config.Budget.TimeBudget)
		defer cancel()
	}

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()
	errGrp, ctx := errgroup.WithContext(ctx)
	for i := 0; i < c.config.Concurrency; i++ {
		errGrp.Go(func() error {
			var result error
			for request := range reqsChan {
				err := c.executor.Execute(request)
				if err != nil && !c.config.ContinueOnError {
					cancel()
					result = multierror.Append(result, err)
				}
			}
			return result
		})
	}

	errGrp.Go(func() error {
		defer close(reqsChan)
		reqs, err := c.source.Requests(ctx)
		if err != nil {
			return err
		}

		c.logger.Info("%d requests to run", len(reqs))

		for _, req := range reqs {
			select {
			case <-ctx.Done():
				return nil
			default:
				transformed, err := c.applyTransformations(req)
				if err != nil {
					return err
				}
				reqsChan <- transformed
			}
		}
		return nil
	})


	return errGrp.Wait()
}

func (c Controller) applyTransformations(req requests.Request) (requests.Request, error) {
	var err error = nil
	for _, transf := range c.transformations {
		req, err = transf.Apply(req)
		if err != nil {
			return requests.Request{}, err
		}
	}

	return req, nil
}
