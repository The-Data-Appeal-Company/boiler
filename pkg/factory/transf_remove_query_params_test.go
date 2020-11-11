package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/transformation"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateRemoveQueryParamsTransformConfiguration(t *testing.T) {
	config, err := createTransformationRemoveQueryParamsConfiguration(conf.TransformationModel{
		Type: transformation.TransformRemoveQueryParams,
		Params: map[string]interface{}{
			"fields": []string{"from", "to"},
		},
	})

	require.NoError(t, err)
	require.Equal(t, config.Fields, []string{"from", "to"})
}
