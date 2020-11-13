package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/controller"
	"boiler/pkg/logging"
	"fmt"
)

func CreateExecutor(model conf.RequestExecutorModel, logger logging.Logger) (controller.Executor, error) {
	switch model.Type {
	case controller.ExecutorHttp:
		return createHttpExecutor(model, logger)
	default:
		return nil, fmt.Errorf("no executor found for type: %s", model.Type)
	}
}
