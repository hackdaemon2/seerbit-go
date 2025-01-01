package payment

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/util"
)

type PaymentEngineProcessor interface {
	ProcessPayment(any, string, string, client.AuthType) (any, error)
}

type PaymentEngine struct {
	Client *client.SeerBitClient
}

func NewPaymentEngine(client *client.SeerBitClient) *PaymentEngine {
	return &PaymentEngine{Client: client}
}

func (paymentEngine *PaymentEngine) ProcessPayment(
	payload any,
	paymentUrl, successCode string,
	authType client.AuthType) (any, error) {
	if !paymentEngine.Client.IsInitialized() {
		return model.PaymentResponse{}, errors.New(client.INITIALIZATION_ERROR)
	}

	var paymentResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	httpRequest := util.HttpRequestData{
		Payload:       payload,
		Response:      &paymentResponse,
		ErrorResponse: &errorResponse,
		Url:           paymentUrl,
		AuthType:      string(authType),
		PrivateKey:    paymentEngine.Client.PrivateKey,
		PublicKey:     paymentEngine.Client.PublicKey,
	}

	if authType == client.Bearer {
		httpRequest.Authentication = paymentEngine.Client.BearerToken
	}

	resp, err := httpRequest.HttpPost()
	if err != nil {
		return nil, fmt.Errorf(client.ERROR_MESSAGE, err)
	}

	shouldReturn, _, err := httpRequest.IsErrorResponse(resp, &errorResponse, &paymentResponse)
	if shouldReturn && len(errorResponse.Error) != 0 {
		return errorResponse, err
	}

	if resp.StatusCode() == http.StatusOK && paymentResponse.Data.Code == successCode {
		log.Println("payment successfully initiated")
	}

	return paymentResponse, nil
}
