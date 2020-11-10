package requests

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestRequestCreation(t *testing.T) {
	req := Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost:1234",
		Path:   "/test",
		Params: map[string][]string{
			"query": {"test_value_0", "test_value_1"},
		},
		Headers: map[string]string{},
	}

	reqUri := req.Uri()

	require.Equal(t, "http://localhost:1234/test?query=test_value_0&query=test_value_1", reqUri.String())
	require.Equal(t, "GET", req.Method)
}

func TestRequestCreationWithoutPort(t *testing.T) {
	req := Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost",
		Path:   "/test",
		Params: map[string][]string{
			"query": {"test_value_0", "test_value_1"},
		},
		Headers: map[string]string{},
	}

	reqUri := req.Uri()

	require.Equal(t, "http://localhost/test?query=test_value_0&query=test_value_1", reqUri.String())
	require.Equal(t, "GET", req.Method)
}

func TestRequestCreationWithoutParams(t *testing.T) {
	req := Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost",
		Path:   "/test",
	}

	reqUri := req.Uri()

	require.Equal(t, "http://localhost/test", reqUri.String())
	require.Equal(t, "GET", req.Method)
}

func TestRequestCreationWithoutPath(t *testing.T) {
	req := Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost",
	}

	reqUri := req.Uri()

	require.Equal(t, "http://localhost", reqUri.String())
	require.Equal(t, "GET", req.Method)
}

func TestRequestCreationWithoutPathWithParams(t *testing.T) {
	req := Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost",
		Params: map[string][]string{
			"query": {"test_value_0", "test_value_1"},
		},
	}

	reqUri := req.Uri()
	require.Equal(t, "http://localhost?query=test_value_0&query=test_value_1", reqUri.String())
	require.Equal(t, "GET", req.Method)
}

func TestRequestCreationFromUrl(t *testing.T) {
	testUrl, err := url.Parse("http://localhost:4321/test?param1=1&param2=2&param2=3")
	require.NoError(t, err)

	req := FromUrl(testUrl, "GET")

	require.Equal(t, "localhost:4321", req.Host)
	require.Equal(t, "http", req.Scheme)
	require.Equal(t, "/test", req.Path)

	require.Contains(t, req.Params, "param1")
	require.Contains(t, req.Params, "param2")

	require.Equal(t, []string{"1"}, req.Params["param1"])
	require.Equal(t, []string{"2", "3"}, req.Params["param2"])
}
