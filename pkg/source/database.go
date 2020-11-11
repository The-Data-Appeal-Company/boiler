package source

import (
	"boiler/pkg/requests"
	"context"
	"database/sql"
	"fmt"
	"net/url"
)

const (
	SourceDatabase = "database"
)

type DatabaseSourceConfiguration struct {
	Connection struct {
		Uri    string
		Driver string
	}
	Extraction struct {
		Query                string
		UrlColumnName        string
		HttpMethodColumnName string
	}
}

type Database struct {
	conf DatabaseSourceConfiguration
}

func NewDatabase(conf DatabaseSourceConfiguration) Database {
	return Database{conf: conf}
}

func (d Database) Requests(ctx context.Context, f func(requests.Request) error) error {
	db, err := sql.Open(d.conf.Connection.Driver, d.conf.Connection.Uri)
	if err != nil {
		return err
	}

	defer db.Close()

	rows, err := db.QueryContext(ctx, d.conf.Extraction.Query)
	if err != nil {
		return err
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	var pointers = make([]interface{}, len(columns))
	for i := range pointers {
		var ii interface{}
		pointers[i] = &ii
	}

	for rows.Next() {
		// We may need to create a slice of pointers to interface{} here
		// better add some unit tests
		if err := rows.Scan(pointers...); err != nil {
			return err
		}

		var namedValues = make(map[string]interface{})
		for i, v := range pointers {
			columnName := columns[i] // this can't be out of bound since we are using rows columns
			namedValues[columnName] = *(v.(*interface{}))
		}

		req, err := d.createRequest(namedValues)
		if err != nil {
			return err
		}

		if err := f(req); err != nil {
			return err
		}
	}

	err = rows.Close()
	if err != nil {
		return err
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (d Database) createRequest(values map[string]interface{}) (requests.Request, error) {
	rawUrl, err := getString(values, d.conf.Extraction.UrlColumnName)
	if err != nil {
		return requests.Request{}, err
	}

	httpMethod, err := getString(values, d.conf.Extraction.HttpMethodColumnName)
	if err != nil {
		return requests.Request{}, err
	}

	uri, err := url.Parse(rawUrl)
	if err != nil {
		return requests.Request{}, err
	}

	req, err := requests.FromUrl(uri, httpMethod)
	if err != nil {
		return requests.Request{}, err
	}
	req.SourceParams = values

	return req, nil
}

func getString(values map[string]interface{}, key string) (string, error) {
	rawUrl, present := values[key]
	if !present {
		return "", fmt.Errorf("database url column not present in response row: %s", key)
	}

	rawUrlStr, isStr := rawUrl.([]byte)
	if !isStr {
		return "", fmt.Errorf("row value for column %s must be a string", key)
	}

	return string(rawUrlStr), nil
}
