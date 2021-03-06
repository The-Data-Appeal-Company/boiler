package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/controller"
	"boiler/pkg/logging"
)

func createHttpExecutor(model conf.RequestExecutorModel, logger logging.Logger) (controller.Executor, error) {
	config, err := createHttpExecutorConfig(model)
	if err != nil {
		return nil, err
	}

	return controller.NewHttpExecutor(config, logger), nil
}

func createHttpExecutorConfig(model conf.RequestExecutorModel) (controller.HttpExecutorConfiguration, error) {
	timeout, err := getDuration(model.Params, "timeout")
	if err != nil {
		return controller.HttpExecutorConfiguration{}, err
	}

	return controller.HttpExecutorConfiguration{Timeout: timeout}, nil
}
