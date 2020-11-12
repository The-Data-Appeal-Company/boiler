package source

import (
	"boiler/pkg/requests"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)
import _ "github.com/proullon/ramsql/driver"

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
