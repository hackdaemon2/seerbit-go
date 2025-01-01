package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
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
	AuthType       string
}

// Centralized HTTP request handler
func (h *HttpRequestData) httpRequest(method string) (*resty.Response, error) {
	client := h.prepareClient()
	var resp *resty.Response
	var err error

	switch method {
	case http.MethodPost:
		resp, err = client.Post(h.Url)
	case http.MethodPut:
		resp, err = client.Put(h.Url)
	case http.MethodGet:
		resp, err = client.Get(h.Url)
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return nil, err
	}

	shouldReturn, result, err := unmarshalResponse(resp, h)
	if shouldReturn {
		return result, err
	}

	return resp, err
}

func (h *HttpRequestData) HttpPost() (*resty.Response, error) {
	return h.httpRequest(http.MethodPost)
}

func (h *HttpRequestData) HttpPut() (*resty.Response, error) {
	return h.httpRequest(http.MethodPut)
}

func (h *HttpRequestData) HttpGet() (*resty.Response, error) {
	return h.httpRequest(http.MethodGet)
}

// Unified response unmarshal function
func unmarshalResponse(resp *resty.Response, h *HttpRequestData) (bool, *resty.Response, error) {
	if err := json.Unmarshal(resp.Body(), h.Response); err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		if err := json.Unmarshal(resp.Body(), h.ErrorResponse); err != nil {
			log.Printf("Error unmarshaling error response: %v", err)
			return true, nil, err
		}
		return true, nil, fmt.Errorf("error response detected")
	}
	return false, nil, nil
}

// Prepares a configured Resty client
func (h *HttpRequestData) prepareClient() *resty.Request {
	client := resty.
		New().
		R().
		SetResult(h.Response).
		SetError(h.ErrorResponse)

	if h.Payload != nil {
		client.SetBody(h.Payload)
	}

	h.setAuthentication(client)
	return client
}

// Adds authentication to the request
func (h *HttpRequestData) setAuthentication(client *resty.Request) {
	switch h.AuthType {
	case "Bearer":
		client.SetAuthToken(h.Authentication)
	case "Basic":
		client.SetBasicAuth(h.PublicKey, h.PrivateKey)
	}
}

// Determines if the response indicates an error
func (h *HttpRequestData) IsErrorResponse(
	resp *resty.Response,
	errResponse *model.ErrorResponse,
	paymentResponse *model.PaymentResponse) (bool, any, error) {

	if resp.StatusCode() != http.StatusOK {
		if err := json.Unmarshal(resp.Body(), errResponse); err != nil {
			return true, nil, fmt.Errorf("failed to parse error response: %w", err)
		}
		return true, *errResponse, nil
	}

	if len(paymentResponse.Error) > 0 {
		return true, &model.ErrorResponse{
			Message: paymentResponse.Message,
			Error:   paymentResponse.Error,
		}, nil
	}

	return false, nil, nil
}
