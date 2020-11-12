package controller

import (
	"boiler/pkg/requests"
	"io/ioutil"
	"net/http"
	"sync"
)

type MockServer struct {
	ServerRequests []requests.Request
	mutex          *sync.Mutex
}

func (m *MockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	req, done := m.buildRequest(r)
	if !done {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m.ServerRequests = append(m.ServerRequests, req)
	if req.Path == "/error" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *MockServer) buildRequest(r *http.Request) (requests.Request, bool) {
	req, err := requests.FromUrl(r.URL, r.Method)
	if err != nil {
		return requests.Request{}, false
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return requests.Request{}, false
	}
	req.Body = string(body)
	req.Headers = r.Header
	return req, true
}

func NewMockServer() *MockServer {
	return &MockServer{
		ServerRequests: make([]requests.Request, 0),
		mutex:          &sync.Mutex{},
	}
}
