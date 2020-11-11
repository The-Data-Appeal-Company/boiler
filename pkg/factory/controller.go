package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/controller"
	"boiler/pkg/transformation"
)

func CreateController(config conf.Config) (controller.Controller, error) {
	source, err := CreateSource(config.Source)
	if err != nil {
		return controller.Controller{}, err
	}

	transformations := make([]transformation.Transformation, len(config.Transformations))
	for i, tModel := range config.Transformations {
		transf, err := CreateTransformation(tModel)
		if err != nil {
			return controller.Controller{}, err
		}

		transformations[i] = transf
	}

	requestExecutor, err := CreateExecutor(config.RequestExecutorModel)
	if err != nil {
		return controller.Controller{}, err
	}

	return controller.NewController(requestExecutor, source, transformations), nil
}