package transformation

import (
	"boiler/pkg/requests"
)

const TransformRemoveQueryParams = "RemoveFilters"

type RemoveQueryParamsTransformConfiguration struct {
	Fields []string
}

type RemoveQueryParamsTransform struct {
	config RemoveQueryParamsTransformConfiguration
}

func NewRemoveFilters(config RemoveQueryParamsTransformConfiguration) RemoveQueryParamsTransform {
	return RemoveQueryParamsTransform{
		config: config,
	}
}

func (r RemoveQueryParamsTransform) Apply(request requests.Request) (requests.Request, error) {
	for _, k := range r.config.Fields {
		delete(request.Params, k)
	}

	return request, nil
}
