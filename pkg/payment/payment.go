package payment

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/constant"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/util"
)

type PaymentEngine struct {
	Client *client.SeerBitClient
}

func NewPaymentEngine(client *client.SeerBitClient) *PaymentEngine {
	return &PaymentEngine{Client: client}
}

func (paymentEngine *PaymentEngine) ProcessPayment(
	payload any,
	paymentUrl, successCode string,
	authType constant.AuthType) (any, error) {
	if !paymentEngine.Client.IsInitialized() {
		return model.PaymentResponse{}, errors.New(constant.INITIALIZATION_ERROR)
	}

	var paymentResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	httpRequest := util.HttpRequestData{
		Payload:       payload,
		Response:      paymentResponse,
		ErrorResponse: errorResponse,
		Url:           paymentUrl,
		AuthType:      authType,
		PrivateKey:    paymentEngine.Client.PrivateKey,
		PublicKey:     paymentEngine.Client.PublicKey,
	}

	if authType == constant.Bearer {
		httpRequest.Authentication = paymentEngine.Client.BearerToken
	}

	resp, err := httpRequest.HttpPost()
	if err != nil {
		return model.PaymentResponse{}, fmt.Errorf(constant.ERROR_MESSAGE, err)
	}

	shouldReturn, paymentError, err := util.IsErrorResponse(resp, errorResponse)
	if shouldReturn {
		return paymentError.(model.ErrorResponse), err
	}

	if resp.StatusCode() == http.StatusOK && paymentResponse.Data.Code == successCode {
		log.Println("payment successfully initiated")
	}

	return paymentResponse, nil
}
