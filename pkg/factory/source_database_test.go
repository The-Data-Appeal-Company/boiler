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
