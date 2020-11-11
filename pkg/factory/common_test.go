package factory

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCommonGetMapStringString(t *testing.T) {
	var data = map[string]interface{}{
		"test": map[string]string{
			"test": "test",
		},
	}

	val, err := getMapStringString(data, "test")
	require.NoError(t, err)
	require.Equal(t, map[string]string{
		"test": "test",
	}, val)
}

func TestCommonGetDuration(t *testing.T) {
	var data = map[string]interface{}{
		"test": "3s",
	}

	val, err := getDuration(data, "test")
	require.NoError(t, err)
	require.Equal(t, 3*time.Second, val)
}

func TestCommonGetBool(t *testing.T) {
	var data = map[string]interface{}{
		"test": true,
	}

	val, err := getBool(data, "test")
	require.NoError(t, err)
	require.Equal(t, true, val)
}

func TestCommonGetStringArrayWithoutType(t *testing.T) {
	var data = map[string]interface{}{
		"test": []interface{}{"1", "2", "3"},
	}

	val, err := getStringArray(data, "test")
	require.NoError(t, err)
	require.Equal(t, []string{"1", "2", "3"}, val)
}

func TestCommonGetStringArray(t *testing.T) {
	var data = map[string]interface{}{
		"test": []string{"1", "2", "3"},
	}

	val, err := getStringArray(data, "test")
	require.NoError(t, err)
	require.Equal(t, []string{"1", "2", "3"}, val)
}

func TestCommonGetInt(t *testing.T) {
	var data = map[string]interface{}{
		"test": 3,
	}

	val, err := getInt(data, "test")
	require.NoError(t, err)
	require.Equal(t, 3, val)
}

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
