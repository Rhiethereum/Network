package common

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type HttpRequestData struct {
	method  string
	uri     string
	body    []byte
	headers map[string]string
}

func (reqData *HttpRequestData) SetMethod(method string) *HttpRequestData {
	return reqData
}

func (reqData *HttpRequestData) SetURI(uri string) *HttpRequestData {
	reqData.uri = uri
	return reqData
}

func (reqData *HttpRequestData) SetHeader(key string, value string) *HttpRequestData {
	reqData.headers[key] = value
	return reqData
}

func (reqData *HttpRequestData) SetBody(body interface{}) error {

	var bodyBytes []byte
	var err error

	bodyBytes, err = json.Marshal(body)
	if err != nil {
		return err
	}

	reqData.body = bodyBytes
	return nil
}

func HttpRequest(reqData HttpRequestData) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(reqData.method)
	req.Header.SetRequestURI(reqData.uri)

	for k, v := range reqData.headers {
		req.Header.Set(k, v)
	}

	if reqData.body != nil {
		req.Header.SetContentLength(len(reqData.body))
		req.SetBody(reqData.body)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
