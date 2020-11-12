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


/*
func TestHttpRequestExecutor_StartExecuteStop(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	mockServer := NewMockServer()
	srv := httptest.NewServer(mockServer)
	defer srv.Close()
	executor := NewHttpRequestExecutor(HttpExecutorConfig{
		Timeout:         1 * time.Second,
		Concurrency:     3,
		ContinueOnError: true,
	})

	reqs := make(chan requests.Request)
	err := executor.Execute(ctx, cancel, reqs)
	require.NoError(t, err)

	nReqs := 10

	for i := 0; i < nReqs; i++ {
		reqs <- getRequest(t, srv.URL)
	}
	close(reqs)
	err = executor.Stop()
	require.NoError(t, err)

	require.Len(t, mockServer.ServerRequests, nReqs)
}

func getRequest(t *testing.T, url string) requests.Request {
	req, err := requests.FromStr(url, "GET")
	require.NoError(t, err)
	return req
}

func TestHttpRequestExecutor_ShouldErrorIfExecutedTwoTimesInARow(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	executor := NewHttpRequestExecutor(HttpExecutorConfig{
		Timeout:         1 * time.Second,
		Concurrency:     3,
		ContinueOnError: true,
	})
	reqs := make(chan requests.Request)
	err := executor.Execute(ctx, cancel, reqs)
	require.NoError(t, err)
	err = executor.Execute(ctx, cancel, reqs)
	require.Error(t, err)
	close(reqs)
	executor.Stop()
}

func TestHttpRequestExecutor_ShouldNoErrorWhenExecuteStopExecute(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	executor := NewHttpRequestExecutor(HttpExecutorConfig{
		Timeout:         1 * time.Second,
		Concurrency:     3,
		ContinueOnError: true,
	})
	reqs := make(chan requests.Request)
	err := executor.Execute(ctx, cancel, reqs)
	require.NoError(t, err)
	close(reqs)
	err = executor.Stop()
	require.NoError(t, err)
	err = executor.Execute(ctx, cancel, reqs)
	require.NoError(t, err)
	err = executor.Stop()
	require.NoError(t, err)
}


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
