package order

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/util"
)

const PRODUCT_ORDERS_ENDPOINT = "/products/orders"

type Order struct {
	Client *client.SeerBitClient
}

func NewOrder(client *client.SeerBitClient) *Order {
	return &Order{
		Client: client,
	}
}

func (order *Order) Create(orderPayload model.OrderPaymentPayload) (any, error) {
	url := order.Client.BaseUrl + "/payments/order"
	return order.executeRequest(orderPayload, url, http.MethodPost)
}

func (order *Order) PostCreate(orderPayload model.OrderPostPaymentPayload) (any, error) {
	url := order.Client.BaseUrl + PRODUCT_ORDERS_ENDPOINT
	return order.executeRequest(orderPayload, url, http.MethodPost)
}

func (order *Order) GetOrders() (any, error) {
	url := strings.Join([]string{order.Client.BaseUrl, PRODUCT_ORDERS_ENDPOINT, "/publicKey/", order.Client.PublicKey}, "")
	return order.executeRequest(nil, url, http.MethodGet)
}

// API reference https://seerbit.github.io/openapi/#operation/GetOrderByPublicKey
func (order *Order) GetOrdersByPaymentReference(paymentReference string) (any, error) {
	url := getOrderReferenceUrl(order, "/paymentReference/", paymentReference)
	return order.executeRequest(nil, url, http.MethodGet)
}

func (order *Order) GetOrdersByOrderId(orderId string) (any, error) {
	url := getOrderReferenceUrl(order, "/orderId/", orderId)
	return order.executeRequest(nil, url, http.MethodGet)
}

func (order *Order) Update(orderPayload model.OrderPostPaymentPayload) (any, error) {
	url := order.Client.BaseUrl + PRODUCT_ORDERS_ENDPOINT
	return order.executeRequest(orderPayload, url, http.MethodPut)
}

func getOrderReferenceUrl(order *Order, referenceUrlPath, referenceValue string) string {
	return strings.Join([]string{
		order.Client.BaseUrl,
		PRODUCT_ORDERS_ENDPOINT,
		"/publicKey/",
		order.Client.PublicKey,
		referenceUrlPath,
		referenceValue,
	}, "")
}

// Common logic for creating and sending HTTP requests
func (order *Order) executeRequest(orderPayload any, url, method string) (any, error) {
	if !order.Client.IsInitialized() {
		return nil, errors.New(client.INITIALIZATION_ERROR)
	}

	var orderResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	httpRequest := util.HttpRequestData{
		PrivateKey:     order.Client.PrivateKey,
		PublicKey:      order.Client.PublicKey,
		Payload:        orderPayload,
		Response:       orderResponse,
		ErrorResponse:  errorResponse,
		Url:            url,
		Authentication: order.Client.BearerToken,
		AuthType:       string(client.Bearer),
	}

	var resp *resty.Response
	var err error

	switch method {
	case http.MethodPost:
		resp, err = httpRequest.HttpPost()
	case http.MethodPut:
		resp, err = httpRequest.HttpPut()
	case http.MethodGet:
		resp, err = httpRequest.HttpGet()
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	shouldReturn, _, err := httpRequest.IsErrorResponse(resp, &errorResponse, &orderResponse)
	if shouldReturn && len(errorResponse.Error) != 0 {
		return errorResponse, err
	}

	return orderResponse, nil
}
