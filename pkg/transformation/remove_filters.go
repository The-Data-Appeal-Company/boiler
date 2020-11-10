package transformation

import (
	"boiler/pkg/requests"
)

const TransformRemoveFilters = "RemoveFilters"

type RemoveFiltersTransform struct {
}

func NewRemoveFilters() RemoveFiltersTransform {
	return RemoveFiltersTransform{}
}

func (r RemoveFiltersTransform) Apply(request requests.Request) (requests.Request, error) {
	request.Params = nil
	return request, nil
}
