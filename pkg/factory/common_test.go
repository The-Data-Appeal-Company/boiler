package factory

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommonGetString(t *testing.T) {
	var data = map[string]interface{}{
		"test": "ok",
	}

	val, err := getString(data, "test")
	require.NoError(t, err)
	require.Equal(t, "ok", val)
}

func TestCommonGetStringNotPresent(t *testing.T) {
	var data = map[string]interface{}{
		"test": "ok",
	}

	_, err := getString(data, "not-present")
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrKeyNotPresent))
}

func TestCommonGetStringInvalidType(t *testing.T) {
	var data = map[string]interface{}{
		"test": false,
	}

	_, err := getString(data, "test")
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrInvalidType))
}
