package controller

import (
	"boiler/pkg/logging"
	"boiler/pkg/requests"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const ExecutorHttp = "http"

type HttpExecutor struct {
	client *http.Client
	logger logging.Logger
}

type HttpExecutorConfiguration struct {
	Timeout time.Duration
}

func NewHttpExecutor(conf HttpExecutorConfiguration, logger logging.Logger) *HttpExecutor {
	return &HttpExecutor{
		logger: logger,
		client: &http.Client{
			Timeout: conf.Timeout,
		},
	}
}

func (f *HttpExecutor) Execute(request requests.Request) error {
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
	resp, err := f.client.Do(httpReq)

	if err != nil {
		return err
	}

	f.logger.Info("%d - %s", resp.StatusCode, request.Uri().String())

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http call status: %s", resp.Status)
	}
	return nil
}
