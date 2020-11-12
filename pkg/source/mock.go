package source

import (
	"boiler/pkg/requests"
	"context"
)

type MockSource struct {
	reqs []requests.Request
}

func NewMockSource(reqs ...requests.Request) *MockSource {
	return &MockSource{reqs: reqs}
}

func (m *MockSource) Requests(ctx context.Context)  ([]requests.Request, error) {
	return m.reqs, nil
}
