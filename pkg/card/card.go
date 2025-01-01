package card

import (
	"errors"
	"fmt"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/payment"
	"github.com/hackdaemon2/seerbit-go/util"
)

// INVALID_PAYLOAD_ERROR represents an error message for invalid payloads.
const INVALID_PAYLOAD_ERROR = "invalid payload"

// Card handles card-based payment operations using the SeerBit API.
type Card struct {
	Client        *client.SeerBitClient  // The SeerBit client for API communication.
	PaymentEngine *payment.PaymentEngine // Payment engine for processing payments.
}

// NewCard initializes a new instance of the Card struct.
//
// Parameters:
//   - client (*client.SeerBitClient): The SeerBit client used for API communication.
//
// Returns:
//   - *Card: A new instance of the Card struct.
func NewCard(client *client.SeerBitClient) *Card {
	return &Card{
		Client:        client,
		PaymentEngine: payment.NewPaymentEngine(client),
	}
}

// Pay initiates a card payment.
//
// Parameters:
//   - payload (any): The payment details. Must be of type model.CardPayload.
//
// Returns:
//   - (any): The response from the SeerBit API.
//   - (error): An error if the operation fails.
//
// For more details see https://seerbit.github.io/openapi/#operation/InitiatePayment
func (card *Card) Pay(payload any) (any, error) {
	cardPayment, ok := payload.(model.CardPayload)
	if !ok {
		return nil, errors.New(INVALID_PAYLOAD_ERROR)
	}
	paymentUrl := card.Client.BaseUrl + "/payments/initiates"
	return card.PaymentEngine.ProcessPayment(cardPayment, paymentUrl, client.SEERBIT_PENDING_CODE, client.Bearer)
}

// CreateToken creates a payment token using the provided card details.
//
// Parameters:
//   - payload (any): The payment details. Must be of type model.CardPayload.
//
// Returns:
//   - (any): The response from the SeerBit API.
//   - (error): An error if the operation fails.
//
// For more details see https://doc.seerbit.com/online-payment/payment-features/card-tokenisation
func (card *Card) CreateToken(payload any) (any, error) {
	cardPayment, ok := payload.(model.CardPayload)
	if !ok {
		return nil, errors.New(INVALID_PAYLOAD_ERROR)
	}
	paymentUrl := card.Client.BaseUrl + "/payments/create-token"
	return card.PaymentEngine.ProcessPayment(cardPayment, paymentUrl, client.SEERBIT_PENDING_CODE, client.Bearer)
}

// ChargeToken charges a tokenized payment.
//
// Parameters:
//   - payload (any): The token charge details. Must be of type model.ChargeToken.
//
// Returns:
//   - (any): The response from the SeerBit API.
//   - (error): An error if the operation fails.
//
// For more details see https://doc.seerbit.com/online-payment/payment-features/card-tokenisation
func (card *Card) ChargeToken(payload any) (any, error) {
	cardPayment, ok := payload.(model.ChargeToken)
	if !ok {
		return nil, errors.New(INVALID_PAYLOAD_ERROR)
	}
	paymentUrl := card.Client.BaseUrl + "/payments/charge-token"
	return card.PaymentEngine.ProcessPayment(cardPayment, paymentUrl, client.SEERBIT_PENDING_CODE, client.Bearer)
}

// DirectCharge initiates a direct charge payment.
//
// Parameters:
//   - payload (model.CardPayload): The payment details.
//
// Returns:
//   - (any): The response from the SeerBit API.
//   - (error): An error if the operation fails.
//
// For more details see https://seerbit.github.io/openapi/#operation/ChargeCard
func (card *Card) DirectCharge(payload model.CardPayload) (any, error) {
	return card.preauthorizeCharge(payload, "/payments/charge")
}

// Authorize initiates a preauthorization for a payment.
//
// Parameters:
//   - payload (model.CardPayload): The payment details.
//
// Returns:
//   - (any): The response from the SeerBit API.
//   - (error): An error if the operation fails.
//
// For more details see https://seerbit.github.io/openapi/#tag/AUTHORISE
func (card *Card) Authorize(payload model.CardPayload) (any, error) {
	return card.preauthorizeCharge(payload, "/payments/authorise")
}

// Authorize3DS initiates a 3D Secure preauthorization for a payment.
//
// Parameters:
//   - payload (model.CardPayload): The payment details.
//
// Returns:
//   - (any): The response from the SeerBit API.
//   - (error): An error if the operation fails.
//
// For more details see https://seerbit.github.io/openapi/#operation/3DSAuthorise
func (card *Card) Authorize3DS(payload model.CardPayload) (any, error) {
	paymentUrl := card.Client.BaseUrl + "/payments/authorise3ds"
	return card.PaymentEngine.ProcessPayment(payload, paymentUrl, client.SEERBIT_PENDING_CODE, client.Bearer)
}

// DirectTokenize tokenizes a card directly without creating a payment.
//
// Parameters:
//   - payload (model.CardPayload): The card details to be tokenized.
//
// Returns:
//   - (any): The tokenization response.
//   - (error): An error if the operation fails.
//
// For more details see https://seerbit.github.io/openapi/#tag/TOKENIZE
func (card *Card) DirectTokenize(payload model.CardPayload) (any, error) {
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
		AuthType:      string(client.Basic),
	}

	resp, err := httpRequest.HttpPost()
	if err != nil {
		return nil, fmt.Errorf(client.ERROR_MESSAGE, err)
	}

	shouldReturn, _, err := httpRequest.IsErrorResponse(resp, &errorResponse, &tokenizationResponse)
	if shouldReturn && len(errorResponse.Error) != 0 {
		return errorResponse, err
	}

	return tokenizationResponse, nil
}

// Capture processes the capture of a previously authorized payment.
//
// Parameters:
//   - payload (model.PreauthorizationPayload): The payload containing preauthorization details such as
//     transaction references and amounts.
//
// Returns:
//   - (any): The response from the SeerBit API for the capture operation.
//   - (error): An error if the operation fails or the API returns an error.
//
// For more details see https://seerbit.github.io/openapi/#operation/Capture
func (card *Card) Capture(payload model.PreauthorizationPayload) (any, error) {
	return card.preauthorizeBackoffice(payload, "/payments/capture")
}

// Refund initiates a refund for a previously completed payment.
//
// Parameters:
//   - payload (model.PreauthorizationPayload): The payload containing details of the transaction to be refunded,
//     including the transaction reference and refund amount.
//
// Returns:
//   - (any): The response from the SeerBit API for the refund operation.
//   - (error): An error if the operation fails or the API returns an error.
//
// For more details see https://seerbit.github.io/openapi/#operation/Refund
func (card *Card) Refund(payload model.PreauthorizationPayload) (any, error) {
	return card.preauthorizeBackoffice(payload, "/payments/refund")
}

// Cancel cancels a previously authorized payment.
//
// Parameters:
//   - payload (model.PreauthorizationPayload): The payload containing details of the authorization to cancel,
//     including the transaction reference.
//
// Returns:
//   - (any): The response from the SeerBit API for the cancel operation.
//   - (error): An error if the operation fails or the API returns an error.
//
// For more details see https://seerbit.github.io/openapi/#operation/Cancel
func (card *Card) Cancel(payload model.PreauthorizationPayload) (any, error) {
	return card.preauthorizeBackoffice(payload, "/payments/cancel")
}

// preauthorizeCharge calls pre-authorize charge endpoints.
//
// Parameters:
//   - payload (model.CardPayload): The payload containing details of the card,
//     including the transaction reference.
//   - endpoint (string): Endpoint to call.
//
// Returns:
//   - (any): The response from the SeerBit API for the cancel operation.
//   - (error): An error if the operation fails or the API returns an error.
func (card *Card) preauthorizeCharge(payload model.CardPayload, endpoint string) (any, error) {
	directChargeUrl := card.Client.BaseUrl + endpoint
	return card.PaymentEngine.ProcessPayment(payload, directChargeUrl, client.SEERBIT_SUCCESS_CODE, client.Basic)
}

// preauthorizeBackoffice calls pre-authorize back-office endpoints.
//
// Parameters:
//   - payload (model.PreauthorizationPayload): The payload containing details of the transaction,
//     including the transaction reference.
//   - endpoint (string): Endpoint to call.
//
// Returns:
//   - (any): The response from the SeerBit API for the cancel operation.
//   - (error): An error if the operation fails or the API returns an error.
func (card *Card) preauthorizeBackoffice(payload model.PreauthorizationPayload, endpoint string) (any, error) {
	directChargeUrl := card.Client.BaseUrl + endpoint
	return card.PaymentEngine.ProcessPayment(payload, directChargeUrl, client.SEERBIT_SUCCESS_CODE, client.Basic)
}
