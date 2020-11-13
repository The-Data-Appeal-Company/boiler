package controller

import (
	"boiler/pkg/requests"
	"fmt"
	"strings"
	"sync"
	"time"
)

type MockExecutorOption interface {
	Execute(req requests.Request) error
}

type MockExecutorDelay struct {
	t time.Duration
}

func MockDelay(t time.Duration) MockExecutorOption {
	return MockExecutorDelay{t: t}
}

func (m MockExecutorDelay) Execute(req requests.Request) error {
	time.Sleep(m.t)
	return nil
}

type MockRequestExecutor struct {
	opts     []MockExecutorOption
	requests []requests.Request
	mutex    *sync.Mutex
}

func NewMockRequestExecutor(opts ...MockExecutorOption) *MockRequestExecutor {
	return &MockRequestExecutor{
		opts:     opts,
		mutex:    &sync.Mutex{},
		requests: make([]requests.Request, 0),
	}
}

func (m *MockRequestExecutor) Execute(req requests.Request) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, opt := range m.opts {
		if err := opt.Execute(req); err != nil {
			return err
		}
	}

	m.requests = append(m.requests, req)
	if strings.Contains(req.Path, "/error") {
		return fmt.Errorf("error")
	}
	return nil
}
