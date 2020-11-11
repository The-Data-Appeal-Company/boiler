package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/source"
)

func createSourceDatabase(model conf.SourceModel) (source.Source, error) {
	configuration, err := createSourceDatabaseConfiguration(model)
	if err != nil {
		return nil, err
	}

	return source.NewDatabase(configuration), nil
}

func createSourceDatabaseConfiguration(model conf.SourceModel) (source.DatabaseSourceConfiguration, error) {
	query, err := getString(model.Params, "query")
	if err != nil {
		return source.DatabaseSourceConfiguration{}, err
	}

	connUri, err := getString(model.Params, "uri")
	if err != nil {
		return source.DatabaseSourceConfiguration{}, err
	}

	connDriver, err := getString(model.Params, "driver")
	if err != nil {
		return source.DatabaseSourceConfiguration{}, err
	}

	urlColumn, err := getString(model.Params, "url_column")
	if err != nil {
		return source.DatabaseSourceConfiguration{}, err
	}

	httpMethodColumn, err := getString(model.Params, "http_method_column")
	if err != nil {
		return source.DatabaseSourceConfiguration{}, err
	}

	return source.DatabaseSourceConfiguration{
		Connection: struct {
			Uri    string
			Driver string
		}{
			Uri:    connUri,
			Driver: connDriver,
		},
		Extraction: struct {
			Query                string
			UrlColumnName        string
			HttpMethodColumnName string
		}{
			Query:                query,
			UrlColumnName:        urlColumn,
			HttpMethodColumnName: httpMethodColumn,
		},
	}, nil
}
