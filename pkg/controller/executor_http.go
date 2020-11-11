package controller

import (
	"boiler/pkg/requests"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type HttpExecutorConfig struct {
	Timeout         time.Duration
	Concurrency     int
	ContinueOnError bool
}

type HttpRequestExecutor struct {
	client *http.Client
	config HttpExecutorConfig
}

func NewHttpRequestExecutor(config HttpExecutorConfig) *HttpRequestExecutor {
	return &HttpRequestExecutor{
		config: config,

		client: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

func (h *HttpRequestExecutor) Execute(request requests.Request) error {
	uri := request.Uri().String()
	var err error
	var resp *http.Response
	switch request.Method {
	case requests.GET:
		resp, err = h.client.Get(uri)
	case requests.POST:
		resp, err = h.client.Post(uri, request.Body.ContentType, strings.NewReader(request.Body.Content))
	}

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("server response is %s: ", resp.Status)
	}
	return nil
}
