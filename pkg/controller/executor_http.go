package controller

import (
	"boiler/pkg/requests"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

const RequestExecutorHttp string = "http"

type HttpExecutorConfig struct {
	Timeout         time.Duration
	Concurrency     int
	ContinueOnError bool
}

type HttpRequestExecutor struct {
	config          HttpExecutorConfig
	requestsChannel chan requests.Request
	errGrp          *errgroup.Group
	executing       bool
}

func NewHttpRequestExecutor(config HttpExecutorConfig) *HttpRequestExecutor {
	return &HttpRequestExecutor{
		config:    config,
		executing: false,
	}
}

func (h *HttpRequestExecutor) Feed(request requests.Request) {
	h.requestsChannel <- request
}

func (h *HttpRequestExecutor) Execute(ctx context.Context) error {
	if h.executing {
		return fmt.Errorf("alredy in execution. please stop before executing again")
	}
	h.executing = true
	h.requestsChannel = make(chan requests.Request, h.config.Concurrency)
	h.errGrp, ctx = errgroup.WithContext(ctx)
	for i := 0; i < h.config.Concurrency; i++ {
		worker := NewFastHttpWorker(h.config.Timeout)
		h.errGrp.Go(func() error {
			for request := range h.requestsChannel {
				fmt.Println("EXE req")
				err := worker.Work(request)
				if err != nil && !h.config.ContinueOnError {
					return err
				}
			}
			return nil
		})
	}
	return nil
}

func (h *HttpRequestExecutor) Stop() error {
	close(h.requestsChannel)
	err := h.errGrp.Wait()
	h.executing = false
	return err
}
