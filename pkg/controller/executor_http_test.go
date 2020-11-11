package controller

import (
	"boiler/pkg/requests"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHttpExecutor(t *testing.T) {
	var serverRequests = make([]requests.Request, 0)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := requests.FromUrl(r.URL, r.Method)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		serverRequests = append(serverRequests, req)

		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	executor := NewHttpRequestExecutor(HttpExecutorConfig{
		Timeout: 5 * time.Second,
	})

	req, err := requests.FromStr(srv.URL, "GET")
	require.NoError(t, err)

	err = executor.Execute(req)
	require.NoError(t, err)

	require.Len(t, serverRequests, 1)
	// TODO Create more assertions
}
