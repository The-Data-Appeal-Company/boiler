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
	begin := time.Now()

	httpReq := f.httpRequest(request)
	resp, err := f.client.Do(httpReq)

	if err != nil {
		f.logger.Warn("%s - %s", request.Uri().String(), err.Error())
		return err
	}

	latency := time.Now().Sub(begin)
	f.logger.Info("%d ms - [%d] - %s", latency.Milliseconds(), resp.StatusCode, request.Uri().Path)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http call status: %s", resp.Status)
	}

	return nil
}

func (f *HttpExecutor) httpRequest(request requests.Request) *http.Request {
	httpReq := &http.Request{
		URL:    request.Uri(),
		Header: request.Headers,
	}

	methodRaw, err := request.Method.String()
	if err != nil {
		f.logger.Warn("unsupported http method: %s", err.Error())
	}

	httpReq.Method = methodRaw

	if f.supportBody(request.Method) {
		httpReq.Body = ioutil.NopCloser(strings.NewReader(request.Body))
	}

	return httpReq
}

func (f *HttpExecutor) supportBody(method requests.HttpMethod) bool {
	return method == "POST"
}
