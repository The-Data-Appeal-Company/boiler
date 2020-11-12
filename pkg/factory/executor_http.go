package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/controller"
)

func createHttpExecutor(model conf.RequestExecutorModel) (controller.Executor, error) {
	config, err := createHttpExecutorConfig(model)
	if err != nil {
		return nil, err
	}

	return controller.NewHttpExecutor(config), nil
}

func createHttpExecutorConfig(model conf.RequestExecutorModel) (controller.HttpExecutorConfiguration, error) {
	timeout, err := getDuration(model.Params, "timeout")
	if err != nil {
		return controller.HttpExecutorConfiguration{}, err
	}

	return controller.HttpExecutorConfiguration{Timeout: timeout}, nil
}
