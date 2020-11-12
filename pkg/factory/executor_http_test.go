package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/controller"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestExecutorHttpCreation(t *testing.T) {
	config, err := createHttpExecutorConfig(conf.RequestExecutorModel{
		Type: controller.ExecutorHttp,
		Params: map[string]interface{}{
			"timeout": "3s",
		},
	})

	require.NoError(t, err)

	require.Equal(t, config.Timeout, 3*time.Second)
}

func TestExecutorHttpCreationInvalidTimeout(t *testing.T) {
	_, err := createHttpExecutorConfig(conf.RequestExecutorModel{
		Type: controller.ExecutorHttp,
		Params: map[string]interface{}{
			"timeout": "3years",
		},
	})

	require.Error(t, err)
}
