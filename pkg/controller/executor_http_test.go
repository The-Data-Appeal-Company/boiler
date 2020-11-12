package controller

import (
	"boiler/pkg/requests"
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHttpRequestExecutor_StartExecuteStop(t *testing.T) {
	mockServer := NewMockServer()
	srv := httptest.NewServer(mockServer)
	defer srv.Close()
	executor := NewHttpRequestExecutor(HttpExecutorConfig{
		Timeout:         1 * time.Second,
		Concurrency:     3,
		ContinueOnError: true,
	})

	err := executor.Execute(context.TODO())
	require.NoError(t, err)

	nReqs := 10
	for i := 0; i < nReqs; i++ {
		executor.Feed(getRequest(t, srv.URL))
	}

	executor.Stop()

	require.Len(t, mockServer.ServerRequests, nReqs)
}

func getRequest(t *testing.T, url string) requests.Request {
	req, err := requests.FromStr(url, "GET")
	require.NoError(t, err)
	return req
}

func TestHttpRequestExecutor_ShouldErrorIfExecutedTwoTimesInARow(t *testing.T) {
	executor := NewHttpRequestExecutor(HttpExecutorConfig{
		Timeout:         1 * time.Second,
		Concurrency:     3,
		ContinueOnError: true,
	})
	err := executor.Execute(context.TODO())
	require.NoError(t, err)
	err = executor.Execute(context.TODO())
	require.Error(t, err)
	executor.Stop()
}

func TestHttpRequestExecutor_ShouldNoErrorWhenExecuteStopExecute(t *testing.T) {
	executor := NewHttpRequestExecutor(HttpExecutorConfig{
		Timeout:         1 * time.Second,
		Concurrency:     3,
		ContinueOnError: true,
	})
	err := executor.Execute(context.TODO())
	require.NoError(t, err)
	executor.Stop()
	err = executor.Execute(context.TODO())
	require.NoError(t, err)
	executor.Stop()
}

/*
func TestHttpRequestExecutor_ContinueOnErrorFalse(t *testing.T) {
	nConcurrency := 3
	mockServer := NewMockServer()
	srv := httptest.NewServer(mockServer)
	defer srv.Close()
	executor := NewHttpRequestExecutor(HttpExecutorConfig{
		Timeout:         1 * time.Second,
		Concurrency:     nConcurrency,
		ContinueOnError: false,
	})

	err := executor.Execute(context.TODO())
	require.NoError(t, err)

	nReqs := 10
	for i := 0; i < nReqs; i++ {
		fmt.Println("MELANI")
		executor.Feed(getRequest(t, srv.URL + "/error"))
	}

	err = executor.Stop()
	require.NoError(t, err)

	require.Len(t, mockServer.ServerRequests, nConcurrency)
}*/
