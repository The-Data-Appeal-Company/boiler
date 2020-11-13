package source

import (
	"boiler/pkg/requests"
	"context"
	"database/sql"
	_ "github.com/proullon/ramsql/driver"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDatabaseSource(t *testing.T) {
	const dbConn = "TestDatabaseSource"

	db, err := sql.Open("ramsql", dbConn)
	if err != nil {
		t.Fatalf("sql.Open : Error : %s\n", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE urls (uri VARCHAR(256) PRIMARY KEY,  method CHAR(10));`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO urls (uri, method) VALUES (?, ?);`, "http://localhost:4321/test?param=1&param=2", "GET")
	require.NoError(t, err)

	config := DatabaseSourceConfiguration{
		Connection: struct {
			Uri    string
			Driver string
		}{
			Uri:    dbConn,
			Driver: "ramsql",
		},
		Extraction: struct {
			Query                string
			UrlColumnName        string
			HttpMethodColumnName string
		}{
			Query:                `select * from urls`,
			UrlColumnName:        "uri",
			HttpMethodColumnName: "method",
		},
	}

	source := NewDatabase(config)

	reqs, err := source.Requests(context.TODO())
	require.NoError(t, err)

	require.Len(t, reqs, 1)

	require.Equal(t, reqs[0], requests.Request{
		Method: "GET",
		Scheme: "http",
		Host:   "localhost:4321",
		Path:   "/test",
		Params: map[string][]string{
			"param": {"1", "2"},
		},
		SourceParams: map[string]interface{}{
			"uri":    []byte("http://localhost:4321/test?param=1&param=2"),
			"method": []byte("GET"),
		},
		Headers: nil,
	})

	uri := reqs[0].Uri()

	require.Equal(t, "http://localhost:4321/test?param=1&param=2", uri.String())
}

func Test_getString(t *testing.T) {
	type args struct {
		values map[string]interface{}
		key    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "shouldErrorWhenKeyNotPresent",
			args: args{
				values: map[string]interface{}{
					"a": "val",
				},
				key: "key",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "shouldErrorWhenTypeNotValid",
			args: args{
				values: map[string]interface{}{
					"a": 15,
				},
				key: "a",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "shouldExtractString",
			args: args{
				values: map[string]interface{}{
					"a": "val",
				},
				key: "a",
			},
			want:    "val",
			wantErr: false,
		},
		{
			name: "shouldExtractBytes",
			args: args{
				values: map[string]interface{}{
					"a": []byte("val"),
				},
				key: "a",
			},
			want:    "val",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getString(tt.args.values, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_createRequest(t *testing.T) {
	type fields struct {
		conf DatabaseSourceConfiguration
	}
	type args struct {
		values map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    requests.Request
		wantErr bool
	}{
		{
			name: "shouldErrorWhenUrlColumnNameNotFound",
			fields: fields{
				conf: DatabaseSourceConfiguration{
					Extraction: struct {
						Query                string
						UrlColumnName        string
						HttpMethodColumnName string
					}{
						UrlColumnName: "url_column",
					},
				},
			},
			args: args{
				values: map[string]interface{}{
					"url": "val",
				},
			},
			want:    requests.Request{},
			wantErr: true,
		},
		{
			name: "shouldErrorWhenHttpMethodColumnNameNotFound",
			fields: fields{
				conf: DatabaseSourceConfiguration{
					Extraction: struct {
						Query                string
						UrlColumnName        string
						HttpMethodColumnName string
					}{
						UrlColumnName:        "url",
						HttpMethodColumnName: "b",
					},
				},
			},
			args: args{
				values: map[string]interface{}{
					"url": "val",
				},
			},
			want:    requests.Request{},
			wantErr: true,
		},
		{
			name: "shouldErrorWhenInvalidUrl",
			fields: fields{
				conf: DatabaseSourceConfiguration{
					Extraction: struct {
						Query                string
						UrlColumnName        string
						HttpMethodColumnName string
					}{
						UrlColumnName:        "url",
						HttpMethodColumnName: "method",
					},
				},
			},
			args: args{
				values: map[string]interface{}{
					"url":    "http://invalid url.com",
					"method": "get",
				},
			},
			want:    requests.Request{},
			wantErr: true,
		},
		{
			name: "shouldErrorWhenInvalidMethod",
			fields: fields{
				conf: DatabaseSourceConfiguration{
					Extraction: struct {
						Query                string
						UrlColumnName        string
						HttpMethodColumnName string
					}{
						UrlColumnName:        "url",
						HttpMethodColumnName: "method",
					},
				},
			},
			args: args{
				values: map[string]interface{}{
					"url":    "http://url.com",
					"method": "delete",
				},
			},
			want:    requests.Request{},
			wantErr: true,
		},
		{
			name: "shouldCreateRequest",
			fields: fields{
				conf: DatabaseSourceConfiguration{
					Extraction: struct {
						Query                string
						UrlColumnName        string
						HttpMethodColumnName string
					}{
						UrlColumnName:        "url",
						HttpMethodColumnName: "method",
					},
				},
			},
			args: args{
				values: map[string]interface{}{
					"url":    "http://url.com/arg",
					"method": "get",
				},
			},
			want: requests.Request{
				Method:       "GET",
				Host:         "url.com",
				Path:         "/arg",
				Scheme:       "http",
				Params:       map[string][]string{},
				SourceParams: map[string]interface{}{"method": "get", "url": "http://url.com/arg"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Database{
				conf: tt.fields.conf,
			}
			got, err := d.createRequest(tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("createRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
