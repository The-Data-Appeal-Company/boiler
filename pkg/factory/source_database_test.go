package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/source"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateSourceDatabase(t *testing.T) {
	uri := "localhost:5432"
	driver := "pg"
	query := `select 1`
	urlColumnName := "url"
	httpMethodName := "http_method"

	config, err := createSourceDatabaseConfiguration(conf.SourceModel{
		Type: "database",
		Params: map[string]interface{}{
			"uri":                uri,
			"driver":             driver,
			"query":              query,
			"url_column":         urlColumnName,
			"http_method_column": httpMethodName,
		},
	})

	require.NoError(t, err)
	require.Equal(t, config, source.DatabaseSourceConfiguration{
		Connection: struct {
			Uri    string
			Driver string
		}{
			Uri:    uri,
			Driver: driver,
		},
		Extraction: struct {
			Query                string
			UrlColumnName        string
			HttpMethodColumnName string
		}{
			Query:                query,
			UrlColumnName:        urlColumnName,
			HttpMethodColumnName: httpMethodName,
		},
	})
}

func Test_createSourceDatabase(t *testing.T) {
	type args struct {
		model conf.SourceModel
	}
	tests := []struct {
		name    string
		args    args
		want    source.Source
		wantErr bool
	}{
		{
			name:    "shouldErrorWhenNoQuery",
			args:    args{
				model: conf.SourceModel{
					Type:   "",
					Params: map[string]interface{}{
						"uri": "http://uri.com",
						"driver": "mysql",
						"url_column": "url",
						"http_method_column": "get",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "shouldErrorWhenNoUri",
			args:    args{
				model: conf.SourceModel{
					Type:   "",
					Params: map[string]interface{}{
						"query": "select * from aa",
						"driver": "mysql",
						"url_column": "url",
						"http_method_column": "get",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "shouldErrorWhenNoDriver",
			args:    args{
				model: conf.SourceModel{
					Type:   "",
					Params: map[string]interface{}{
						"query": "select * from aa",
						"uri": "http://uri.com",
						"url_column": "url",
						"http_method_column": "get",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "shouldErrorWhenNoUrlColumn",
			args:    args{
				model: conf.SourceModel{
					Type:   "",
					Params: map[string]interface{}{
						"query": "select * from aa",
						"uri": "http://uri.com",
						"driver": "mysql",
						"http_method_column": "get",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "shouldErrorWhenNoMethod",
			args:    args{
				model: conf.SourceModel{
					Type:   "",
					Params: map[string]interface{}{
						"query": "select * from aa",
						"uri": "http://uri.com",
						"driver": "mysql",
						"url_column": "url",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createSourceDatabase(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("createSourceDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}