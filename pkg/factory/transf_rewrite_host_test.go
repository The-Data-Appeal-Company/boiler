package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/transformation"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateRewriteHostTransformConfiguration(t *testing.T) {
	const host = "127.0.0.1"

	config, err := createTransformationRewriteHostConfiguration(conf.TransformationModel{
		Type: transformation.TransformRewriteHost,
		Params: map[string]interface{}{
			"host": host,
		},
	})

	require.NoError(t, err)
	require.Equal(t, config.Host, host)
}
