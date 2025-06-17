package test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func Add(requestPath string, statusCode int, body interface{}) {
	interceptor := &Interceptor{}
	interceptor.Add(requestPath, statusCode, body)
}

type Interceptor struct {
	Transport http.RoundTripper
	option    Option
}

type Option struct {
	requestPath string
	statusCode  int
	body        interface{}
}

func (interceptor *Interceptor) Add(requestPath string, statusCode int, body interface{}) {
	interceptor.option = Option{
		requestPath: requestPath,
		statusCode:  statusCode,
		body:        body,
	}
	interceptor.Transport = http.DefaultTransport
	http.DefaultClient = &http.Client{Transport: interceptor}
}

func (interceptor *Interceptor) RoundTrip(req *http.Request) (*http.Response, error) {

	validResponse, _ := json.Marshal(interceptor.option.body)

	return &http.Response{
		StatusCode: interceptor.option.statusCode,
		Body:       io.NopCloser(strings.NewReader(string(validResponse))),
	}, nil

}
