package requests

import (
	"fmt"
	"net/url"
)

type Request struct {
	Method       string
	Scheme       string
	Host         string
	Path         string
	Params       map[string][]string
	Headers      map[string]string
	SourceParams map[string]interface{}
}

func FromUrl(u *url.URL, method string) Request {
	return Request{
		Method: method,
		Scheme: u.Scheme,
		Host:   u.Host,
		Path:   u.Path,
		Params: u.Query(),
	}
}

func (r Request) Uri() *url.URL {
	u := &url.URL{
		Scheme: r.Scheme,
		Host:   r.Host,
		Path:   r.Path,
	}

	q := u.Query()
	for key, param := range r.Params {
		for _, value := range param {
			q.Add(key, value)
		}
	}

	u.RawQuery = q.Encode()

	return u
}

func (r Request) String() interface{} {
	return fmt.Sprintf("%s %s", r.Method, r.Uri().String())
}
