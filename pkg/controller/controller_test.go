package controller

import (
	"boiler/pkg/logging"
	"boiler/pkg/requests"
	"boiler/pkg/source"
	"boiler/pkg/transformation"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var testLogger = logging.Noop()

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
	}, testLogger)

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
	}, testLogger)

	err = contrl.Execute(context.TODO())
	require.Error(t, err)
	require.GreaterOrEqual(t, len(requestExecutor.requests), 1)
}

func TestHttpExecutor_TimeBudget(t *testing.T) {
	testWithTimeout(t, 10*time.Second, func(t *testing.T) {
		requestExecutor := NewMockRequestExecutor(
			MockDelay(100 * time.Millisecond),
		)

		req, err := requests.FromStr("http//localhost:4321/", "GET")
		require.NoError(t, err)

		reqsCount := 1000
		reqs := make([]requests.Request, reqsCount)
		for i := 0; i < reqsCount; i++ {
			reqs[i] = req
		}

		src := source.NewMockSource(reqs...)

		contrl := NewController(src, []transformation.Transformation{}, requestExecutor, Config{
			Concurrency:     3,
			ContinueOnError: false,
			Budget:          BudgetConfig{TimeBudget: 300 * time.Millisecond},
		}, testLogger)

		err = contrl.Execute(context.TODO())
		require.Error(t, context.DeadlineExceeded)
	})
}

func TestHttpExecutor_TimeBudgetCompletion(t *testing.T) {
	testWithTimeout(t, 10*time.Second, func(t *testing.T) {
		requestExecutor := NewMockRequestExecutor(MockDelay(2 * time.Millisecond))

		req, err := requests.FromStr("http//localhost:4321/", "GET")
		require.NoError(t, err)

		reqsCount := 1000
		reqs := make([]requests.Request, reqsCount)
		for i := 0; i < reqsCount; i++ {
			reqs[i] = req
		}

		src := source.NewMockSource(reqs...)

		contrl := NewController(src, []transformation.Transformation{}, requestExecutor, Config{
			Concurrency:     10,
			ContinueOnError: false,
			Budget:          BudgetConfig{TimeBudget: 1 * time.Hour},
		}, testLogger)

		err = contrl.Execute(context.TODO())
		require.NoError(t, err)

		require.Len(t, requestExecutor.requests, reqsCount)
	})
}

func testWithTimeout(t *testing.T, timeoutDuration time.Duration, testFn func(t *testing.T)) {
	done := make(chan interface{})
	timeout := time.After(timeoutDuration)

	go func() {
		testFn(t)
		done <- true
	}()

	select {
	case <-done:
		break
	case <-timeout:
		t.Fatal("test timed out")
	}
}
