package factory

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrKeyNotPresent = errors.New("key not present")
	ErrInvalidType   = errors.New("invalid key value type")
)

func getDuration(m map[string]interface{}, key string) (time.Duration, error) {
	rawValue, present := m[key]
	if !present {
		return 0, fmt.Errorf("%w: %s", ErrKeyNotPresent, key)
	}

	value, isStr := rawValue.(string)
	if !isStr {
		return 0, fmt.Errorf("%w: %s", ErrInvalidType, key)
	}

	durationValue, err := time.ParseDuration(value)
	if err != nil {
		return 0, err
	}

	return durationValue, nil
}

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

func getInt(m map[string]interface{}, key string) (int, error) {
	rawValue, present := m[key]
	if !present {
		return 0, fmt.Errorf("%w: %s", ErrKeyNotPresent, key)
	}

	value, isInt := rawValue.(int)
	if !isInt {
		return 0, fmt.Errorf("%w: %s", ErrInvalidType, key)
	}

	return value, nil
}

func getBool(m map[string]interface{}, key string) (bool, error) {
	rawValue, present := m[key]
	if !present {
		return false, fmt.Errorf("%w: %s", ErrKeyNotPresent, key)
	}

	value, isBool := rawValue.(bool)
	if !isBool {
		return false, fmt.Errorf("%w: %s", ErrInvalidType, key)
	}

	return value, nil
}

func getStringArray(m map[string]interface{}, key string) ([]string, error) {
	rawValue, present := m[key]
	if !present {
		return nil, fmt.Errorf("%w: %s", ErrKeyNotPresent, key)
	}

	// if we already have a string array we can just cast and return
	valueStr, isStrArr := rawValue.([]string)
	if isStrArr {
		return valueStr, nil
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
