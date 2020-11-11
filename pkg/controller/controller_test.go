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
	req, err := requests.FromStr("http//localhost:4321?test=true", "GET")

	requestExecutor := NewMockRequestExecutor()

	src := source.NewMockSource(req, req, req)

	transformations := []transformation.Transformation{
		transformation.NewRemoveFilters(),
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
	}

}
