package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/transformation"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateWriteHostTransformConfiguration(t *testing.T) {
	headers := map[string]string{
		"test": "true",
	}

	config, err := createTransformationWriteHeaderConfiguration(conf.TransformationModel{
		Type: transformation.TransformRewriteHost,
		Params: map[string]interface{}{
			"headers": headers,
		},
	})

	require.NoError(t, err)
	require.Equal(t, config.Headers, headers)
}
