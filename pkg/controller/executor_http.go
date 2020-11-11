package controller

import (
	"boiler/pkg/requests"
	"net/http"
	"time"
)

const RequestExecutorHttp string = "http"

type HttpExecutorConfig struct {
	Timeout         time.Duration
	Concurrency     int
	ContinueOnError bool
}

type HttpRequestExecutor struct {
	client *http.Client
	config HttpExecutorConfig
}

func NewHttpRequestExecutor(config HttpExecutorConfig) HttpRequestExecutor {
	return HttpRequestExecutor{
		config: config,

		client: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

func (h HttpRequestExecutor) Execute(request requests.Request) error {
	panic("implement me")
}
