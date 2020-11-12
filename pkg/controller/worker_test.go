package controller

import (
	"boiler/pkg/requests"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFastHttpWorker_Work(t *testing.T) {
	type fields struct {
		timeout time.Duration
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
				timeout: 1 * time.Second,
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
			name: "should execute GET",
			fields: fields{
				timeout: 1 * time.Second,
			},
			args: args{
				method: "GET",
				headers: map[string][]string{
					"Language": {
						"it",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "should error when status != 200",
			fields: fields{
				timeout: 1 * time.Second,
			},
			args: args{
				method: "GET",
				path:   "/error",
				headers: map[string][]string{
					"Language": {
						"it",
					},
				},
			},
			wantErr: true,
		},
	}
	mockServer := NewMockServer()
	srv := httptest.NewServer(mockServer)
	defer srv.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer.ServerRequests = make([]requests.Request, 0)
			req, err := requests.FromStr(srv.URL+tt.args.path, tt.args.method)
			require.NoError(t, err)
			req.Headers = tt.args.headers
			req.Body = tt.args.body
			f := NewFastHttpWorker(tt.fields.timeout)
			if err := f.Work(req); (err != nil) != tt.wantErr {
				t.Errorf("Work() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				require.Len(t, mockServer.ServerRequests, 1)
				for k, values := range tt.args.headers {
					for _, value := range values {
						require.Contains(t, mockServer.ServerRequests[0].Headers[k], value)
					}
				}
				require.Equal(t, tt.args.body, mockServer.ServerRequests[0].Body)
			}
		})
	}
}
