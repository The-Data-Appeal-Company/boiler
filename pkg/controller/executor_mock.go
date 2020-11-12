package controller

import (
	"boiler/pkg/requests"
	"fmt"
	"strings"
	"sync"
)

type MockRequestExecutor struct {
	requests []requests.Request
	mutex    *sync.Mutex
}

func NewMockRequestExecutor() *MockRequestExecutor {
	return &MockRequestExecutor{requests: make([]requests.Request, 0), mutex: &sync.Mutex{}}
}

func (m *MockRequestExecutor) Execute(req requests.Request) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.requests = append(m.requests, req)
	if strings.Contains(req.Path, "/error") {
		return fmt.Errorf("error")
	}
	return nil
}
