package controller

import (
	"boiler/pkg/requests"
	"context"
)

type RequestExecutor interface {
	Feed(request requests.Request)
	Execute(ctx context.Context) error
	Stop() error
}
