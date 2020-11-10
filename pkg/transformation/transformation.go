package transformation

import "boiler/pkg/requests"

type Transformation interface {
	Apply(request requests.Request) (requests.Request, error)
}
