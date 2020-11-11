package transformation

import (
	"boiler/pkg/requests"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveFiltersTransform(t *testing.T) {
	req := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost",
		Params: map[string][]string{
			"query": {"test_value_0", "test_value_1"},
		},
	}

	transformed, err := NewRemoveFilters(RemoveQueryParamsTransformConfiguration{
		Fields: []string{"query"},
	}).Apply(req)
	require.NoError(t, err)

	require.Equal(t, req.Host, transformed.Host)
	require.Equal(t, req.Scheme, transformed.Scheme)
	require.Equal(t, req.Method, transformed.Method)

	require.Empty(t, transformed.Params)
}
