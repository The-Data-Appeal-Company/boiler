package factory

import (
	"boiler/pkg/conf"
	"boiler/pkg/controller"
	"boiler/pkg/logging"
	"boiler/pkg/source"
	"boiler/pkg/transformation"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type MockLog struct{}

func (m *MockLog) Debug(s string, i ...interface{}) {}

func (m *MockLog) Info(s string, i ...interface{}) {}

func (m *MockLog) Warn(s string, i ...interface{}) {}

func (m *MockLog) Error(s string, i ...interface{}) {}

func TestControllerCreationFromConfiguration(t *testing.T) {
	config, err := conf.NewFileReader("testdata/example-config.yml").ReadConf()
	require.NoError(t, err)

	controllerConfig, err := createControllerConfig(config)
	require.NoError(t, err)

	require.Equal(t, 32*time.Second, controllerConfig.Budget.TimeBudget)
	require.Equal(t, 1, controllerConfig.Concurrency)
	require.Equal(t, true, controllerConfig.ContinueOnError)
}

func TestControllerCreationFromConfigurationError(t *testing.T) {
	config, err := conf.NewFileReader("testdata/example-config-error.yml").ReadConf()
	require.NoError(t, err)

	_, err = createControllerConfig(config)
	require.Error(t, err)
}

func TestCreateController(t *testing.T) {
	s := source.NewDatabase(source.DatabaseSourceConfiguration{
		Connection: struct {
			Uri    string
			Driver string
		}{
			Uri:    "sslmode=require user=<user> password=<password> host=<host> port=<port> dbname=<db_name>",
			Driver: "postgres",
		},
		Extraction: struct {
			Query                string
			UrlColumnName        string
			HttpMethodColumnName string
		}{
			Query:                "select uri, request_date, 'GET' as method FROM table",
			UrlColumnName:        "uri",
			HttpMethodColumnName: "method",
		},
	})
	trans := []transformation.Transformation{
		transformation.NewRewriteHostTransform(transformation.RewriteHostTransformConfiguration{
			Host: "localhost:8080",
		}),
	}
	logger := &MockLog{}
	exe := controller.NewHttpExecutor(controller.HttpExecutorConfiguration{
		Timeout: 1 * time.Minute,
	}, logger)
	confi := controller.Config{
		Concurrency:     1,
		ContinueOnError: true,
		Budget:          controller.BudgetConfig{
			TimeBudget: 32 * time.Second,
		},
	}
	type args struct {
		config conf.Config
		logger logging.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    controller.Controller
		wantErr bool
	}{
		{
			name: "shouldCreateController",
			args: args{
				config: getConfFromFile("testdata/example-config-simple.yml"),
				logger: &MockLog{},
			},
			want:    controller.NewController(s, trans, exe, confi, logger),
			wantErr: false,
		},
		{
			name: "shouldErrorWrongSource",
			args: args{
				config: getConfFromFile("testdata/example-config-wrong-source.yml"),
				logger: &MockLog{},
			},
			want:    controller.Controller{},
			wantErr: true,
		},
		{
			name: "shouldErrorWrongTransformation",
			args: args{
				config: getConfFromFile("testdata/example-config-wrong-trans.yml"),
				logger: &MockLog{},
			},
			want:    controller.Controller{},
			wantErr: true,
		},
		{
			name: "shouldErrorWrongExecutor",
			args: args{
				config: getConfFromFile("testdata/example-config-wrong-exe.yml"),
				logger: &MockLog{},
			},
			want:    controller.Controller{},
			wantErr: true,
		},
		{
			name: "shouldErrorWrongConfig",
			args: args{
				config: getConfFromFile("testdata/example-config-wrong-conf.yml"),
				logger: &MockLog{},
			},
			want:    controller.Controller{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateController(tt.args.config, tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateController() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func getConfFromFile(file string) conf.Config {
	config, _ := conf.NewFileReader(file).ReadConf()
	return config
}
