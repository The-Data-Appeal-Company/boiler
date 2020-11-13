package factory

import (
	"errors"
	"github.com/stretchr/testify/require"
	"reflect"
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

func TestCommonGetDurationErrorNotPresent(t *testing.T) {
	var data = map[string]interface{}{
		"secs": "3s",
	}

	val, err := getDuration(data, "test")
	require.Error(t, err)
	require.Zero(t, val)
}

func TestCommonGetDurationErrorNotString(t *testing.T) {
	var data = map[string]interface{}{
		"secs": 30,
	}

	val, err := getDuration(data, "secs")
	require.Error(t, err)
	require.Zero(t, val)
}

func TestCommonGetBool(t *testing.T) {
	var data = map[string]interface{}{
		"test": true,
	}

	val, err := getBool(data, "test")
	require.NoError(t, err)
	require.Equal(t, true, val)
}

func TestCommonGetBoolErrorWhenNotPresent(t *testing.T) {
	var data = map[string]interface{}{
		"bool": true,
	}

	val, err := getBool(data, "test")
	require.Error(t, err)
	require.False(t, val)
}

func TestCommonGetBoolErrorWhenNotBool(t *testing.T) {
	var data = map[string]interface{}{
		"test": "true",
	}

	val, err := getBool(data, "test")
	require.Error(t, err)
	require.False(t, val)
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

func TestCommonGetStringArrayErrorWhenNotPresent(t *testing.T) {
	var data = map[string]interface{}{
		"arr": []string{"1", "2", "3"},
	}

	val, err := getStringArray(data, "test")
	require.Error(t, err)
	require.Nil(t, val)
}

func TestCommonGetStringArrayErrorWhenNotArray(t *testing.T) {
	var data = map[string]interface{}{
		"test": "1",
	}

	val, err := getStringArray(data, "test")
	require.Error(t, err)
	require.Nil(t, val)
}

func TestCommonGetStringArrayErrorWhenNotStringArray(t *testing.T) {
	var data = map[string]interface{}{
		"test": []interface{}{1, 2},
	}

	val, err := getStringArray(data, "test")
	require.Error(t, err)
	require.Nil(t, val)
}

func TestCommonGetInt(t *testing.T) {
	var data = map[string]interface{}{
		"test": 3,
	}

	val, err := getInt(data, "test")
	require.NoError(t, err)
	require.Equal(t, 3, val)
}

func TestCommonGetIntErrorWhenNotPresent(t *testing.T) {
	var data = map[string]interface{}{
		"int": 3,
	}

	val, err := getInt(data, "test")
	require.Error(t, err)
	require.Zero(t, val)
}

func TestCommonGetIntErrorWhenNotInt(t *testing.T) {
	var data = map[string]interface{}{
		"test": "3",
	}

	val, err := getInt(data, "test")
	require.Error(t, err)
	require.Zero(t, val)
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

func Test_getMapStringString(t *testing.T) {
	type args struct {
		m   map[string]interface{}
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "shouldErrorWhenKeyNotPresent",
			args: args{
				m: map[string]interface{}{
					"k": "val",
				},
				key: "key",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "shouldErrorWhenInvalidType",
			args: args{
				m: map[string]interface{}{
					"key": 10,
				},
				key: "key",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "shouldErrorWhenMapInterfaceNoKeyString",
			args: args{
				m: map[string]interface{}{
					"key": map[interface{}]interface{}{10: "nnn"},
				},
				key: "key",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "shouldErrorWhenMapInterfaceNoValueString",
			args: args{
				m: map[string]interface{}{
					"key": map[interface{}]interface{}{"aaa": 10},
				},
				key: "key",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "shouldGetMapString",
			args: args{
				m: map[string]interface{}{
					"key": map[string]string{"aaa": "bbb"},
				},
				key: "key",
			},
			want:    map[string]string{"aaa": "bbb"},
			wantErr: false,
		},
		{
			name: "shouldGetMapInterface",
			args: args{
				m: map[string]interface{}{
					"key": map[interface{}]interface{}{"aaa": "nnn"},
				},
				key: "key",
			},
			want:    map[string]string{"aaa": "nnn"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMapStringString(tt.args.m, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMapStringString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMapStringString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
