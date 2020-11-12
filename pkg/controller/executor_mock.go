package controller

import (
	"boiler/pkg/requests"
	"context"
)

type MockRequestExecutor struct {
	requests []requests.Request
}

func NewMockRequestExecutor() *MockRequestExecutor {
	return &MockRequestExecutor{requests: make([]requests.Request, 0)}
}

func (m *MockRequestExecutor) Execute(ctx context.Context) error {
	return nil
}

func (m *MockRequestExecutor) Feed(request requests.Request) {
	m.requests = append(m.requests, request)
}
func (m *MockRequestExecutor) Stop() error {
	return nil
}
