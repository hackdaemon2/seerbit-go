package card

import (
	"errors"
	"fmt"
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/constant"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/payment"
	"github.com/hackdaemon2/seerbit-go/util"
)

type Card struct {
	Client        *client.SeerBitClient
	paymentEngine *payment.PaymentEngine
}

func NewCard(client *client.SeerBitClient) *Card {
	return &Card{
		Client:        client,
		paymentEngine: payment.NewPaymentEngine(client),
	}
}

func (card *Card) Pay(payload any) (any, error) {
	cardPayment, ok := payload.(model.CardPayload)
	if !ok {
		return nil, errors.New("invalid payload for Card")
	}
	paymentUrl := card.Client.BaseUrl + "/payments/initiates"
	return card.paymentEngine.ProcessPayment(cardPayment, paymentUrl, constant.SEERBIT_PENDING_CODE, constant.Bearer)
}

func (card *Card) DirectCharge(payload model.CardPayload) (any, error) {
	return card.preauthorizeCharge(payload, "/payments/charge")
}

func (card *Card) Authorize(payload model.CardPayload) (any, error) {
	return card.preauthorizeCharge(payload, "/payments/authorise")
}

func (card *Card) Authorize3DS(payload model.CardPayload) (any, error) {
	paymentUrl := card.Client.BaseUrl + "/payments/authorise3ds"
	return card.paymentEngine.ProcessPayment(payload, paymentUrl, constant.SEERBIT_PENDING_CODE, constant.Bearer)
}

func (card *Card) Tokenize(payload model.CardPayload) (any, error) {
	var tokenizationResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	tokenizationUrl := card.Client.BaseUrl + "/payments/tokenize"
	httpRequest := util.HttpRequestData{
		Payload:       payload,
		Response:      tokenizationResponse,
		ErrorResponse: errorResponse,
		Url:           tokenizationUrl,
		PublicKey:     card.Client.PublicKey,
		PrivateKey:    card.Client.PrivateKey,
		AuthType:      constant.Basic,
	}

	resp, err := httpRequest.HttpPost()
	if err != nil {
		return nil, fmt.Errorf(constant.ERROR_MESSAGE, err)
	}

	shouldReturn, checkoutError, err := httpRequest.IsErrorResponse(resp, errorResponse)
	if shouldReturn {
		return checkoutError, err
	}

	log.Println("checkout payment successfully initiated.")
	return tokenizationResponse, nil
}

func (card *Card) Capture(payload model.PreauthorizationPayload) (any, error) {
	return card.preauthorizeBackoffice(payload, "/payments/capture")
}

func (card *Card) Refund(payload model.PreauthorizationPayload) (any, error) {
	return card.preauthorizeBackoffice(payload, "/payments/refund")
}

func (card *Card) Cancel(payload model.PreauthorizationPayload) (any, error) {
	return card.preauthorizeBackoffice(payload, "/payments/cancel")
}

func (card *Card) RecurringSubscription(payload model.CardPayload) (any, error) {
	paymentUrl := card.Client.BaseUrl + "/recurring/subscribes"
	return card.paymentEngine.ProcessPayment(payload, paymentUrl, constant.SEERBIT_PENDING_CODE, constant.Bearer)
}

func (card *Card) GetSubscription(billingId string) (any, error) {
	var subscriptionResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	subscriptionUrl := fmt.Sprintf("%s/recurring/billingId/%s", card.Client.BaseUrl, billingId)
	httpRequest := util.HttpRequestData{
		Response:       subscriptionResponse,
		ErrorResponse:  errorResponse,
		Url:            subscriptionUrl,
		PublicKey:      card.Client.PublicKey,
		PrivateKey:     card.Client.PrivateKey,
		AuthType:       constant.Bearer,
		Authentication: card.Client.BearerToken,
	}

	resp, err := httpRequest.HttpGet()
	if err != nil {
		return nil, fmt.Errorf(constant.ERROR_MESSAGE, err)
	}

	shouldReturn, subscriptionErrResponse, err := httpRequest.IsErrorResponse(resp, errorResponse)
	if shouldReturn {
		return subscriptionErrResponse, err
	}

	log.Println("subscription list successfully retrieved.")
	return subscriptionResponse, nil
}

func (card *Card) preauthorizeCharge(payload model.CardPayload, endpoint string) (any, error) {
	directChargeUrl := card.Client.BaseUrl + endpoint
	return card.paymentEngine.ProcessPayment(payload, directChargeUrl, constant.SEERBIT_SUCCESS_CODE, constant.Basic)
}

func (card *Card) preauthorizeBackoffice(payload model.PreauthorizationPayload, endpoint string) (any, error) {
	directChargeUrl := card.Client.BaseUrl + endpoint
	return card.paymentEngine.ProcessPayment(payload, directChargeUrl, constant.SEERBIT_SUCCESS_CODE, constant.Basic)
}
