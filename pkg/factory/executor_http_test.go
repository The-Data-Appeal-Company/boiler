package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/controller"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestExecutorCreationFromConfiguration(t *testing.T) {
	config, err := conf.NewFileReader("testdata/example-config.yml").ReadConf()
	require.NoError(t, err)

	executorConf, err := createHttpExecutorConfig(config.RequestExecutorModel)
	require.NoError(t, err)

	require.Equal(t, executorConf.Timeout, 60*time.Second)
}

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
