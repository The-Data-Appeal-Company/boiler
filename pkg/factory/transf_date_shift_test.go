package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/transformation"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateTransformationDateShift(t *testing.T) {
	config, err := createTransformationRelativeDateShiftConfiguration(conf.TransformationModel{
		Type: transformation.TransformationRelativeDateShift,
		Params: map[string]interface{}{
			"relative_to":   "$column",
			"date_format":   "2006-01-02",
			"target_fields": []interface{}{"from", "to"},
		},
	})

	require.NoError(t, err)

	require.Equal(t, config.DateFormat, "2006-01-02")
	require.Equal(t, config.RelativeTo, "$column")
	require.Equal(t, config.TargetFields, []string{"from", "to"})
	require.NotNil(t, config.RelativeTimeFn)
}
