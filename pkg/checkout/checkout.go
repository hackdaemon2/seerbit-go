package standardcheckout

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

type Checkout struct {
	Client *client.SeerBitClient
}

func NewCheckout(client *client.SeerBitClient) *Checkout {
	return &Checkout{
		Client: client,
	}
}

func (checkout *Checkout) Pay(payload any) (any, error) {
	checkoutPayload, ok := payload.(CheckoutPayload)
	if !ok {
		return nil, errors.New("invalid payload for Standard Checkout")
	}

	if !checkout.Client.IsInitialized() {
		return nil, errors.New(constant.INITIALIZATION_ERROR)
	}

	var hashResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	hashUrl := checkout.Client.BaseUrl + "/encrypt/hashs"
	token := checkout.Client.BearerToken

	httpRequest := util.HttpRequestData{
		Payload:        checkoutPayload,
		Response:       hashResponse,
		ErrorResponse:  errorResponse,
		Url:            hashUrl,
		Authentication: token,
		AuthType:       constant.Bearer,
	}

	resp, err := httpRequest.HttpPost()
	if err != nil {
		return nil, fmt.Errorf(constant.ERROR_MESSAGE, err)
	}

	shouldReturn, checkoutError, err := util.IsErrorResponse(resp, errorResponse)
	if shouldReturn {
		return checkoutError, err
	}

	success := resp.StatusCode() == http.StatusOK && hashResponse.Data.Code == constant.SEERBIT_SUCCESS_CODE
	if success {
		log.Println("checkout payload successfully hashed.")
		return checkout.handleSuccessHash(hashResponse, checkoutPayload)
	}

	return hashResponse, nil
}

func (checkout *Checkout) handleSuccessHash(hashResponse model.PaymentResponse, payload CheckoutPayload) (any, error) {
	var checkoutResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	hash := hashResponse.Data.Hash.Hash
	hashedPayload := &CheckoutHashPayload{
		CheckoutPayload: payload,
		Hash:            hash,
		HashType:        constant.SEERBIT_HASH_TYPE,
	}

	paymentUrl := checkout.Client.BaseUrl + "/payments"
	token := checkout.Client.BearerToken

	httpRequest := util.HttpRequestData{
		Payload:        hashedPayload,
		Response:       checkoutResponse,
		ErrorResponse:  errorResponse,
		Url:            paymentUrl,
		Authentication: token,
		AuthType:       constant.Bearer,
	}

	resp, err := httpRequest.HttpPost()
	if err != nil {
		return nil, fmt.Errorf(constant.ERROR_MESSAGE, err)
	}

	shouldReturn, checkoutError, err := util.IsErrorResponse(resp, errorResponse)
	if shouldReturn {
		return checkoutError, err
	}

	log.Println("checkout payment successfully initiated.")
	return checkoutResponse, nil
}
