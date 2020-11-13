package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/controller"
	"boiler/pkg/transformation"
	"time"
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

	controllerConf, err := createControllerConfig(config)
	if err != nil {
		return controller.Controller{}, err
	}

	return controller.NewController(source, transformations, requestExecutor, controllerConf), nil
}

func createControllerConfig(config conf.Config) (controller.Config, error) {
	timeBudget, err := time.ParseDuration(config.RequestExecutorModel.Budget.Time)
	if err != nil {
		return controller.Config{}, err
	}

	return controller.Config{
		Concurrency:     config.RequestExecutorModel.Config.Concurrency,
		ContinueOnError: config.RequestExecutorModel.Config.ContinueOnError,
		Budget: controller.BudgetConfig{
			TimeBudget: timeBudget,
		},
	}, nil
}
