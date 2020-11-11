package controller

import (
	"boiler/pkg/requests"
	"boiler/pkg/source"
	"boiler/pkg/transformation"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestController(t *testing.T) {
	req, err := requests.FromStr("http//localhost:4321/test?test=true", "GET")

	requestExecutor := NewMockRequestExecutor()

	src := source.NewMockSource(req, req, req)

	transformations := []transformation.Transformation{
		transformation.NewRemoveQueryFilters(transformation.RemoveQueryParamsTransformConfiguration{
			Fields: []string{"test"},
		}),
	}

	contrl := Controller{
		source:          src,
		transformations: transformations,
		requestExecutor: requestExecutor,
	}

	err = contrl.Execute(context.TODO())
	require.NoError(t, err)

	require.Len(t, requestExecutor.requests, 3)
	for _, r := range requestExecutor.requests {
		require.Empty(t, r.Params)
		require.Equal(t, r.Host, req.Host)
		require.Equal(t, r.Path, req.Path)
		require.Equal(t, r.Scheme, req.Scheme)
	}

}
