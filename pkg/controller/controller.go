package controller

import (
	"boiler/pkg/requests"
	"boiler/pkg/source"
	"boiler/pkg/transformation"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

type Config struct {
	Concurrency     int
	Timeout         time.Duration
	ContinueOnError bool
}

type Controller struct {
	source          source.Source
	transformations []transformation.Transformation
	config          Config
}

func NewController(source source.Source, transformations []transformation.Transformation, config Config) Controller {
	return Controller{
		source:          source,
		transformations: transformations,
		config:          config,
	}
}

func (c Controller) Execute(parentCtx context.Context) error {
	reqsChan := make(chan requests.Request, c.config.Concurrency)
	ctx, cancel := context.WithCancel(parentCtx)
	errGrp, ctx := errgroup.WithContext(ctx)
	for i := 0; i < c.config.Concurrency; i++ {
		executor := NewHttpExecutor(c.config.Timeout)
		errGrp.Go(func() error {
			for request := range reqsChan {
				fmt.Println("EXE req")
				err := executor.Execute(request)
				if err != nil && !c.config.ContinueOnError {
					cancel()
					return err
				}
			}
			return nil
		})
	}

	errGrp.Go(func() error {
		defer close(reqsChan)
		reqs, err := c.source.Requests(ctx)
		if err != nil {
			return err
		}
		for _, req := range reqs {
			select {
			case <-ctx.Done():
				break
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
