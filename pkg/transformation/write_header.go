package transformation

import (
	"boiler/pkg/requests"
)

const TransformWriteHeader = "write-header"

type WriteHeaderTransformConfiguration struct {
	Headers map[string]string
}

type WriteHeaderTransform struct {
	config WriteHeaderTransformConfiguration
}

func NewWriteHeaderTransform(config WriteHeaderTransformConfiguration) WriteHeaderTransform {
	return WriteHeaderTransform{config: config}
}

func (r WriteHeaderTransform) Apply(request requests.Request) (requests.Request, error) {
	for name, value := range r.config.Headers{
		request.Headers[name] = []string{value}
	}
	return request, nil
}
