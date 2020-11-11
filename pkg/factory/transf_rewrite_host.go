package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/transformation"
)

func createTransformationRewriteHost(model conf.TransformationModel) (transformation.Transformation, error) {
	config, err := createTransformationRewriteHostConfiguration(model)
	if err != nil {
		return nil, err
	}

	return transformation.NewRewriteHostTransform(config), nil
}

func createTransformationRewriteHostConfiguration(model conf.TransformationModel) (transformation.RewriteHostTransformConfiguration, error) {
	host, err := getString(model.Params, "host")
	if err != nil {
		return transformation.RewriteHostTransformConfiguration{}, err
	}

	return transformation.RewriteHostTransformConfiguration{
		Host: host,
	}, nil
}
