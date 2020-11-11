package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/controller"
)

func createHttpExecutor(model conf.RequestExecutorModel) (controller.RequestExecutor, error) {
	config, err := createHttpExecutorConfig(model)
	if err != nil {
		return nil, err
	}

	return controller.NewHttpRequestExecutor(config), nil
}

func createHttpExecutorConfig(model conf.RequestExecutorModel) (controller.HttpExecutorConfig, error) {
	timeout, err := getDuration(model.Params, "timeout")
	if err != nil {
		return controller.HttpExecutorConfig{}, err
	}

	concurrency, err := getInt(model.Params, "concurrency")
	if err != nil {
		return controller.HttpExecutorConfig{}, err
	}

	continueOnErr, err := getBool(model.Params, "continue_on_error")
	if err != nil {
		return controller.HttpExecutorConfig{}, err
	}

	config := controller.HttpExecutorConfig{
		Timeout:         timeout,
		Concurrency:     concurrency,
		ContinueOnError: continueOnErr,
	}
	return config, nil
}
