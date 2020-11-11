package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/controller"
	"fmt"
)

func CreateExecutor(model conf.RequestExecutorModel) (controller.RequestExecutor, error) {
	switch model.Type {
	case controller.RequestExecutorHttp:
		return createHttpExecutor(model)
	default:
		return nil, fmt.Errorf("no executor found for type: %s", model.Type)
	}
}
