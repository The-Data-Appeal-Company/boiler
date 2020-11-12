package controller

import (
	"boiler/pkg/requests"
)

type MockRequestExecutor struct {
	requests []requests.Request
}

func NewMockRequestExecutor() *MockRequestExecutor {
	return &MockRequestExecutor{requests: make([]requests.Request, 0)}
}

func (m *MockRequestExecutor) Execute(req requests.Request) error {
	m.requests = append(m.requests, req)
	return nil
}
