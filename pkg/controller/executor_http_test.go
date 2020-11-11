package controller

import (
	"boiler/pkg/requests"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockServer struct {
	serverRequests []requests.Request
}

func (m *MockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, done := m.buildRequest(r)
	if !done {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m.serverRequests = append(m.serverRequests, req)
	if req.Path == "/error" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MockServer) buildRequest(r *http.Request) (requests.Request, bool) {
	req, err := requests.FromUrl(r.URL, r.Method)
	if err != nil {
		return requests.Request{}, false
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return requests.Request{}, false
	}
	req.Body = string(body)
	req.Headers = r.Header
	return req, true
}

func NewMockServer() *MockServer {
	return &MockServer{serverRequests: make([]requests.Request, 0)}
}

func TestHttpExecutor_GET(t *testing.T) {
	mockServer := NewMockServer()
	srv := httptest.NewServer(mockServer)
	defer srv.Close()

	executor := NewHttpRequestExecutor(HttpExecutorConfig{
		Timeout: 5 * time.Second,
	})

	req, err := requests.FromStr(srv.URL, "GET")
	require.NoError(t, err)

	err = executor.Execute(req)
	require.NoError(t, err)

	require.Len(t, mockServer.serverRequests, 1)
}

func TestHttpRequestExecutor_Execute(t *testing.T) {
	type fields struct {
		config HttpExecutorConfig
	}
	type args struct {
		method  string
		path    string
		body    string
		headers map[string][]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should execute POST",
			fields: fields{
				config: HttpExecutorConfig{
					Timeout:     1 * time.Second,
					Concurrency: 1,
				},
			},
			args: args{
				method: "POST",
				body:   "{\"key\": \"value\"}",
				headers: map[string][]string{
					"Content-Type": {
						"application-json",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when http error and ContinueOnError = false",
			fields: fields{
				config: HttpExecutorConfig{
					Timeout:         1 * time.Second,
					Concurrency:     1,
					ContinueOnError: false,
				},
			},
			args: args{
				path:   "/error",
				method: "GET",
			},
			wantErr: true,
		},
		{
			name: "should NOT return error when http error and ContinueOnError = true",
			fields: fields{
				config: HttpExecutorConfig{
					Timeout:         1 * time.Second,
					Concurrency:     1,
					ContinueOnError: true,
				},
			},
			args: args{
				path:   "/error",
				method: "GET",
			},
			wantErr: false,
		},
	}

	mockServer := NewMockServer()
	srv := httptest.NewServer(mockServer)
	defer srv.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer.serverRequests = make([]requests.Request, 0)
			h := NewHttpRequestExecutor(tt.fields.config)
			req, err := requests.FromStr(srv.URL+tt.args.path, tt.args.method)
			require.NoError(t, err)
			req.Headers = tt.args.headers
			req.Body = tt.args.body
			if err := h.Execute(req); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				require.Len(t, mockServer.serverRequests, 1)
				for k, values := range tt.args.headers {
					for _, value := range values {
						require.Contains(t, mockServer.serverRequests[0].Headers[k], value)
					}
				}
				require.Equal(t, tt.args.body, mockServer.serverRequests[0].Body)
			}
		})
	}
}
