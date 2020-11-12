package controller

import (
	"boiler/pkg/requests"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

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

	fmt.Printf("%d - %s\n", resp.StatusCode, request.Uri().String())
	if resp.StatusCode != 200 {
		return fmt.Errorf("http call status: %s", resp.Status)
	}
	return nil
}

type HttpExecutor struct {
	client *http.Client
}

func NewHttpExecutor(timeout time.Duration) *HttpExecutor {
	return &HttpExecutor{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}
