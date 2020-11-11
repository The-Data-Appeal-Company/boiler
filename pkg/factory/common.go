package factory

import (
	"errors"
	"fmt"
)

var (
	ErrKeyNotPresent = errors.New("key not present")
	ErrInvalidType   = errors.New("invalid key value type")
)

func getString(m map[string]interface{}, key string) (string, error) {
	rawValue, present := m[key]
	if !present {
		return "", fmt.Errorf("%w: %s", ErrKeyNotPresent, key)
	}

	value, isStr := rawValue.(string)
	if !isStr {
		return "", fmt.Errorf("%w: %s", ErrInvalidType, key)
	}

	return value, nil
}

func getStringArray(m map[string]interface{}, key string) ([]string, error) {
	rawValue, present := m[key]
	if !present {
		return nil, fmt.Errorf("%w: %s", ErrKeyNotPresent, key)
	}

	value, isArr := rawValue.([]interface{})
	if !isArr {
		return nil, fmt.Errorf("%w: %s", ErrInvalidType, key)
	}

	strValues := make([]string, len(value))
	for i := range value {
		strVal, isStr := value[i].(string)
		if !isStr {
			return nil, fmt.Errorf("%w: %s", ErrInvalidType, key)
		}
		strValues[i] = strVal
	}

	return strValues, nil
}
