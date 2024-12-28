package util

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/hackdaemon2/seerbit-go/pkg/constant"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
)

type HttpRequestData struct {
	Payload        any
	Response       any
	ErrorResponse  any
	Url            string
	Authentication string
	PublicKey      string
	PrivateKey     string
	AuthType       constant.AuthType
}

func (h *HttpRequestData) HttpPost() (*resty.Response, error) {
	client := h.prepareClient()
	return client.Post(h.Url)
}

func (h *HttpRequestData) HttpPut() (*resty.Response, error) {
	client := h.prepareClient()
	return client.Put(h.Url)
}

func (h *HttpRequestData) HttpGet() (*resty.Response, error) {
	client := h.prepareClient()
	return client.Get(h.Url)
}

func (h *HttpRequestData) prepareClient() *resty.Request {
	client := resty.
		New().
		R().
		SetResult(&h.Response).
		SetError(&h.ErrorResponse)

	if h.Payload != nil {
		client.SetBody(h.Payload)
	}

	h.setAuthentication(client)
	return client
}

func (h *HttpRequestData) setAuthentication(client *resty.Request) {
	if h.AuthType == constant.Bearer {
		client.SetAuthToken(h.Authentication)
	} else if h.AuthType == constant.Basic {
		client.SetBasicAuth(h.PublicKey, h.PrivateKey)
	}
}

func IsErrorResponse(resp *resty.Response, errResponse model.ErrorResponse) (bool, any, error) {
	if resp.StatusCode() != http.StatusOK || len(errResponse.Error) != 0 {
		return true, errResponse, nil
	}
	return false, nil, nil
}
