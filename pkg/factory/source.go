package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/source"
	"fmt"
)

func createSource(model conf.SourceModel) (source.Source, error) {
	switch model.Type {
	case source.SourceDatabase:
		return createSourceDatabase(model)
	default:
		return nil, fmt.Errorf("no source found for type: %s", model.Type)
	}
}
