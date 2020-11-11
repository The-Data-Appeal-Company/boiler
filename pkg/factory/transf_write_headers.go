package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/transformation"
)

func createTransformationWriteHeader(model conf.TransformationModel) (transformation.Transformation, error) {
	config, err := createTransformationWriteHeaderConfiguration(model)
	if err != nil {
		return nil, err
	}

	return transformation.NewWriteHeaderTransform(config), nil
}

func createTransformationWriteHeaderConfiguration(model conf.TransformationModel) (transformation.WriteHeaderTransformConfiguration, error) {
	headers, err := getMapStringString(model.Params, "headers")
	if err != nil {
		return transformation.WriteHeaderTransformConfiguration{}, err
	}

	return transformation.WriteHeaderTransformConfiguration{
		Headers: headers,
	}, nil
}
