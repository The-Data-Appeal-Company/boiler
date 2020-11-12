package source

import (
	"boiler/pkg/requests"
	"context"
)

type Source interface {
	Requests(context.Context) ([]requests.Request, error)
}
