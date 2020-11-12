package controller

import (
	"boiler/pkg/requests"
)

type Executor interface {
	Execute(request requests.Request) error
}
