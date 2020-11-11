package transformation

import (
	"boiler/pkg/requests"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRewriteHeaders(t *testing.T) {
	req := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost",
		Params: map[string][]string{
			"query": {"test_value_0", "test_value_1"},
		},
		Headers: map[string][]string{
			"test0": {"test"},
		},
	}

	transformed, err := NewWriteHeaderTransform(WriteHeaderTransformConfiguration{
		Headers: map[string]string{
			"test0": "test0",
			"test1": "value1",
			"test2": "value2",
		},
	}).Apply(req)

	require.NoError(t, err)

	require.Equal(t, req.Host, transformed.Host)
	require.Equal(t, req.Scheme, transformed.Scheme)
	require.Equal(t, req.Method, transformed.Method)
	require.Equal(t, req.Params, transformed.Params)

	require.Equal(t, req.Headers, map[string][]string{
		"test0": {"test0"},
		"test1": {"value1"},
		"test2": {"value2"},
	})
}
