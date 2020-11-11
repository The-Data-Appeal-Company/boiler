package controller

import "boiler/pkg/requests"

type RequestExecutor interface {
	Execute(request requests.Request) error
}

