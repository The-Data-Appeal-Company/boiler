package controller

import (
	"boiler/pkg/requests"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func NewHttpRequestExecutor(config HttpExecutorConfig) *HttpRequestExecutor {
	return &HttpRequestExecutor{
		config: config,
		client: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

func (h *HttpRequestExecutor) Execute(request requests.Request) error {
	uri := request.Uri()
	httpReq := &http.Request{
		URL:    uri,
		Header: request.Headers,
	}
	switch request.Method {
	case requests.GET:
	case requests.POST:
		httpReq.Method = "POST"
		httpReq.Body = ioutil.NopCloser(strings.NewReader(request.Body))
	}
	resp, err := h.client.Do(httpReq)

	if err != nil {
		if h.config.ContinueOnError {
			return nil
		}
		return err
	}

	if resp.StatusCode != 200 && !h.config.ContinueOnError {
		return fmt.Errorf("http call status: %s", resp.Status)
	}

	fmt.Println(uri.String())
	return nil
}
