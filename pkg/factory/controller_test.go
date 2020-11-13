package factory

import (
	"boiler/pkg/conf"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestControllerCreationFromConfiguration(t *testing.T) {
	config, err := conf.NewFileReader("testdata/example-config.yml").ReadConf()
	require.NoError(t, err)

	controllerConfig, err := createControllerConfig(config)
	require.NoError(t, err)

	require.Equal(t, 32*time.Second, controllerConfig.Budget.TimeBudget)
	require.Equal(t, 1, controllerConfig.Concurrency)
	require.Equal(t, true, controllerConfig.ContinueOnError)
}
