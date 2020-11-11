package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/transformation"
)

func createTransformationRemoveQueryParams(model conf.TransformationModel) (transformation.Transformation, error) {
	config, err := createTransformationRemoveQueryParamsConfiguration(model)
	if err != nil {
		return nil, err
	}

	return transformation.NewRemoveFilters(config), nil
}

func createTransformationRemoveQueryParamsConfiguration(model conf.TransformationModel) (transformation.RemoveQueryParamsTransformConfiguration, error) {
	targetFields, err := getStringArray(model.Params, "fields")
	if err != nil {
		return transformation.RemoveQueryParamsTransformConfiguration{}, err
	}

	config := transformation.RemoveQueryParamsTransformConfiguration{
		Fields: targetFields,
	}
	return config, nil
}
