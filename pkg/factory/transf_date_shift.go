package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/transformation"
	"time"
)

func createTransformationRelativeDateShift(model conf.TransformationModel) (transformation.Transformation, error) {
	config, err := createTransformationRelativeDateShiftConfiguration(model)
	if err != nil {
		return nil, err
	}

	return transformation.NewDateShift(config), nil
}

func createTransformationRelativeDateShiftConfiguration(model conf.TransformationModel) (transformation.RelativeDateShiftConfiguration, error) {
	relativeTo, err := getString(model.Params, "relative_to")
	if err != nil {
		return transformation.RelativeDateShiftConfiguration{}, err
	}

	dateFormat, err := getString(model.Params, "date_format")
	if err != nil {
		return transformation.RelativeDateShiftConfiguration{}, err
	}

	targetFields, err := getStringArray(model.Params, "target_fields")
	if err != nil {
		return transformation.RelativeDateShiftConfiguration{}, err
	}

	relativeTimeFn := func() time.Time {
		return time.Now()
	}

	config := transformation.RelativeDateShiftConfiguration{
		RelativeTo:     relativeTo,
		DateFormat:     dateFormat,
		TargetFields:   targetFields,
		RelativeTimeFn: relativeTimeFn,
	}
	return config, nil
}
