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
		Type: controller.RequestExecutorHttp,
		Params: map[string]interface{}{
			"continue_on_error": true,
			"timeout":           "3s",
			"concurrency":       1,
		},
	})

	require.NoError(t, err)

	require.Equal(t, config.Timeout, 3*time.Second)
	require.Equal(t, config.Concurrency, 1)
	require.Equal(t, config.ContinueOnError, true)
}

func TestExecutorHttpCreationInvalidTimeout(t *testing.T) {
	_, err := createHttpExecutorConfig(conf.RequestExecutorModel{
		Type: controller.RequestExecutorHttp,
		Params: map[string]interface{}{
			"continue_on_error": true,
			"timeout":           "3years",
			"concurrency":       1,
		},
	})

	require.Error(t, err)
}

func TestExecutorHttpCreationMissingConcurrency(t *testing.T) {
	_, err := createHttpExecutorConfig(conf.RequestExecutorModel{
		Type: controller.RequestExecutorHttp,
		Params: map[string]interface{}{
			"continue_on_error": true,
			"timeout":           "3s",
		},
	})

	require.Error(t, err)
}
