package checkout

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/util"
)

// Checkout encapsulates the SeerBit Standard Checkout functionality.
type Checkout struct {
	Client *client.SeerBitClient // The SeerBit client used to interact with the API.
}

// NewCheckout creates a new Checkout instance.
//
// Parameters:
//   - client (*client.SeerBitClient): The SeerBit client instance used to interact with the API.
//
// Returns:
//   - (*Checkout): A new Checkout instance.
func NewCheckout(client *client.SeerBitClient) *Checkout {
	return &Checkout{
		Client: client,
	}
}

// Pay initiates a payment using SeerBit's Standard Checkout.
//
// Parameters:
//   - payload (any): The payment payload, expected to be of type `model.CheckoutPayload`.
//     The payload should include details like amount, currency, customer information, etc.
//
// Returns:
//   - (any): The response from the SeerBit API, which typically includes payment status and additional data.
//   - (error): An error if the payload is invalid, the client is not initialized, or the API request fails.
//
// For more details see https://seerbit.github.io/openapi/#operation/PaymentLinkFromCheckout
func (checkout *Checkout) Pay(payload any) (any, error) {
	checkoutPayload, ok := payload.(model.CheckoutPayload)
	if !ok {
		return nil, errors.New("invalid payload for Standard Checkout")
	}

	if !checkout.Client.IsInitialized() {
		return nil, errors.New(client.INITIALIZATION_ERROR)
	}

	var paymentResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	paymentUrl := checkout.Client.BaseUrl + "/payments"
	token := checkout.Client.BearerToken

	httpRequest := util.HttpRequestData{
		Payload:        checkoutPayload,
		Response:       &paymentResponse,
		ErrorResponse:  &errorResponse,
		Url:            paymentUrl,
		Authentication: token,
		AuthType:       string(client.Bearer),
	}

	resp, err := httpRequest.HttpPost()
	if err != nil {
		return nil, fmt.Errorf(client.ERROR_MESSAGE, err)
	}

	shouldReturn, _, err := httpRequest.IsErrorResponse(resp, &errorResponse, &paymentResponse)
	if shouldReturn && len(errorResponse.Error) != 0 {
		return errorResponse, err
	}

	success := resp.StatusCode() == http.StatusOK && paymentResponse.Data.Code == client.SEERBIT_SUCCESS_CODE
	if success {
		log.Println("checkout payment successfully initiated.")
	}

	return paymentResponse, nil
}
