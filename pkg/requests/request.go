package requests

import (
	"fmt"
	"net/url"
)

type HttpMethod string

const (
	GET  HttpMethod = "GET"
	POST HttpMethod = "POST"
)

type Body struct {
	ContentType string
	Content     string
}

type Request struct {
	Method       HttpMethod
	Scheme       string
	Host         string
	Path         string
	Params       map[string][]string
	Headers      map[string]string
	Body         Body
	SourceParams map[string]interface{}
}

func HttpMethodFromString(method string) (HttpMethod, error) {
	switch method {
	case "GET":
		return GET, nil
	case "POST":
		return POST, nil
	}
	return "", fmt.Errorf("method %s not supported", method)
}

func FromStr(u string, method string) (Request, error) {
	uri, err := url.Parse(u)
	if err != nil {
		return Request{}, err
	}

	return FromUrl(uri, method)
}

func FromUrl(u *url.URL, method string) (Request, error) {
	m, err := HttpMethodFromString(method)
	if err != nil {
		return Request{}, err
	}
	return Request{
		Method: m,
		Scheme: u.Scheme,
		Host:   u.Host,
		Path:   u.Path,
		Params: u.Query(),
	}, nil
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
