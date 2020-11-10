package transformation

import "fmt"

func Create(name string, params map[string]string) (Transformation, error) {
	switch name {
	case TransformRemoveFilters:
		return NewRemoveFilters(), nil
	}

	return nil, fmt.Errorf("no transformation found with name: %s", name)
}
