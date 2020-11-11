package transformation

import (
	"boiler/pkg/requests"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRewriteHost(t *testing.T) {
	req := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost",
		Params: map[string][]string{
			"query": {"test_value_0", "test_value_1"},
		},
	}

	const newHost = "127.0.0.1:8888"
	transformed, err := NewRewriteHostTransform(RewriteHostTransformConfiguration{
		Host: newHost,
	}).Apply(req)
	require.NoError(t, err)

	require.Equal(t, newHost, transformed.Host)
	require.Equal(t, req.Scheme, transformed.Scheme)
	require.Equal(t, req.Method, transformed.Method)
	require.Equal(t, req.Params, transformed.Params)

}
