package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/transformation"
	"fmt"
)

func CreateTransformation(model conf.TransformationModel) (transformation.Transformation, error) {
	switch model.Type {
	case transformation.TransformationRelativeDateShift:
		return createTransformationRelativeDateShift(model)
	case transformation.TransformRemoveQueryParams:
		return createTransformationRemoveQueryParams(model)
	case transformation.TransformRewriteHost:
		return createTransformationRewriteHost(model)
	case transformation.TransformWriteHeader:
		return createTransformationWriteHeader(model)
	default:
		return nil, fmt.Errorf("no transformation found for type: %s", model.Type)
	}
}
