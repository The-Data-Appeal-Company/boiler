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

func (m *MockSource) Requests(ctx context.Context, f func(requests.Request) error) error {
	for _, req := range m.reqs {
		if err := f(req); err != nil {
			return err
		}
	}
	return nil
}
