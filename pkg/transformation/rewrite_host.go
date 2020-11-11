package transformation

import (
	"boiler/pkg/requests"
)

const TransformRewriteHost = "rewrite-host"

type RewriteHostTransformConfiguration struct {
	Host string
}

type RewriteHostTransform struct {
	config RewriteHostTransformConfiguration
}

func NewRewriteHostTransform(config RewriteHostTransformConfiguration) RewriteHostTransform {
	return RewriteHostTransform{config: config}
}

func (r RewriteHostTransform) Apply(request requests.Request) (requests.Request, error) {
	request.Host = r.config.Host
	return request, nil
}
