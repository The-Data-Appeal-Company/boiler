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
	default:
		return nil, fmt.Errorf("no transformation found for type: %s", model.Type)
	}
}
