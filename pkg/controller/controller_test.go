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

	src := source.NewMockSource(req, req, req, req, req, req, req, req)

	transformations := []transformation.Transformation{
		transformation.NewRemoveQueryFilters(transformation.RemoveQueryParamsTransformConfiguration{
			Fields: []string{"test"},
		}),
	}
	contrl := NewController(src, transformations, requestExecutor, Config{
		Concurrency:     3,
		ContinueOnError: true,
	})

	err = contrl.Execute(context.TODO())
	require.NoError(t, err)

	require.Len(t, requestExecutor.requests, 8)
	for _, r := range requestExecutor.requests {
		require.Empty(t, r.Params)
		require.Equal(t, r.Host, req.Host)
		require.Equal(t, r.Path, req.Path)
		require.Equal(t, r.Scheme, req.Scheme)
	}
}

func TestHttpRequestExecutor_ContinueOnErrorFalse(t *testing.T) {
	req, err := requests.FromStr("http//localhost:4321/error", "GET")

	requestExecutor := NewMockRequestExecutor()

	src := source.NewMockSource(req, req, req, req, req, req, req, req)

	transformations := []transformation.Transformation{
		transformation.NewRemoveQueryFilters(transformation.RemoveQueryParamsTransformConfiguration{
			Fields: []string{"test"},
		}),
	}

	contrl := NewController(src, transformations, requestExecutor, Config{
		Concurrency:     3,
		ContinueOnError: false,
	})

	err = contrl.Execute(context.TODO())
	require.Error(t, err)
	require.GreaterOrEqual(t, len(requestExecutor.requests), 1)
}
