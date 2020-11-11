package controller

import (
	"boiler/pkg/requests"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMockRequestExecutor(t *testing.T) {
	executor := NewMockRequestExecutor()

	req1 := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost:4321",
	}
	req2 := requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost:4326",
	}

	err := executor.Execute(req1)
	require.NoError(t, err)

	err = executor.Execute(req2)
	require.NoError(t, err)

	require.Len(t, executor.requests, 2)

	require.Equal(t, executor.requests[0], req1)
	require.Equal(t, executor.requests[1], req2)
}
