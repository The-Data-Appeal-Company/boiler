package controller

import (
	"boiler/pkg/requests"
	"boiler/pkg/source"
	"boiler/pkg/transformation"
	"context"
)

type Controller struct {
	source          source.Source
	transformations []transformation.Transformation
	requestExecutor RequestExecutor
}

func NewController(requestExecutor RequestExecutor, source source.Source, transformations []transformation.Transformation) Controller {
	return Controller{
		requestExecutor: requestExecutor,
		source:          source,
		transformations: transformations,
	}
}

func (c Controller) Execute(ctx context.Context) error {
	return c.source.Requests(ctx, func(request requests.Request) error {
		request, err := c.applyTransformations(request)

		if err != nil {
			return err
		}

		if err := c.requestExecutor.Execute(request); err != nil {
			return err
		}

		return nil
	})
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