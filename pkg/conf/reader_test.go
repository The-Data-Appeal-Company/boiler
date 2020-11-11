package conf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileReader(t *testing.T) {
	conf, err := NewFileReader("testdata/example-00.yml").ReadConf()
	require.NoError(t, err)

	require.Equal(t, "database", conf.Source.Type)

	driver, isStr := conf.Source.Params["driver"].(string)
	require.True(t, isStr)
	require.Equal(t, "redshift", driver)

	uri, isStr := conf.Source.Params["uri"].(string)
	require.True(t, isStr)
	require.Equal(t, "root@localhost:5432", uri)

	require.Len(t, conf.Transformations, 2)

	require.Equal(t, conf.Transformations[0].Type, "remove-filters")
	filters, isArr := conf.Transformations[0].Params["filters"].([]interface{})
	require.True(t, isArr)
	require.Equal(t, []interface{}{"to", "from"}, filters)

	require.Equal(t, conf.Transformations[1].Type, "add-filter")

	name, isStr := conf.Transformations[1].Params["name"].(string)
	require.True(t, isStr)
	require.Equal(t, "test", name)

	value, isStr := conf.Transformations[1].Params["value"].(string)
	require.True(t, isStr)
	require.Equal(t, "value", value)
}
