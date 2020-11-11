package transformation

import (
	"boiler/pkg/requests"
)

const TransformRemoveQueryParams = "remove-query-params"

type RemoveQueryParamsTransformConfiguration struct {
	Fields []string
}

type RemoveQueryParamsTransform struct {
	config RemoveQueryParamsTransformConfiguration
}

func NewRemoveQueryFilters(config RemoveQueryParamsTransformConfiguration) RemoveQueryParamsTransform {
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
