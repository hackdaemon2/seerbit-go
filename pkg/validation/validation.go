package validation

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

type Validation struct {
	Client *client.SeerBitClient
}

func NewValidation(client *client.SeerBitClient) *Validation {
	return &Validation{
		Client: client,
	}
}

func (validation *Validation) Verify(reference string) (any, error) {
	shouldReturn, _, err := isInitialized(validation)
	if shouldReturn {
		return nil, err
	}

	var verifyResponse model.PaymentResponse
	var verifyErrorResponse model.ErrorResponse

	validationUrl := fmt.Sprintf("%s/%s/%s", validation.Client.BaseUrl, "payments/query", reference)
	token := validation.Client.BearerToken

	httpRequest := util.HttpRequestData{
		Response:       verifyResponse,
		ErrorResponse:  verifyErrorResponse,
		Url:            validationUrl,
		Authentication: token,
		AuthType:       constant.Bearer,
	}

	resp, err := httpRequest.HttpGet()
	if err != nil {
		return nil, fmt.Errorf(constant.ERROR_MESSAGE, err)
	}

	shouldReturn, verifyErr, err := httpRequest.IsErrorResponse(resp, verifyErrorResponse)
	if shouldReturn {
		return verifyErr, err
	}

	return verifyResponse, nil
}

func (validation *Validation) Validate(payload model.ValidationPayload) (any, error) {
	shouldReturn, _, err := isInitialized(validation)
	if shouldReturn {
		return nil, err
	}

	var validationResponse model.PaymentResponse
	var validationErrorResponse model.ErrorResponse

	validationUrl := validation.Client.BaseUrl + "/payments/validate"
	token := validation.Client.BearerToken

	httpRequest := util.HttpRequestData{
		Payload:        payload,
		Response:       validationResponse,
		ErrorResponse:  validationErrorResponse,
		Url:            validationUrl,
		Authentication: token,
		AuthType:       constant.Bearer,
	}

	resp, err := httpRequest.HttpPost()
	if err != nil {
		return nil, fmt.Errorf(constant.ERROR_MESSAGE, err)
	}

	shouldReturn, validationErr, err := httpRequest.IsErrorResponse(resp, validationErrorResponse)
	if shouldReturn {
		return validationErr, err
	}

	pending := resp.StatusCode() == http.StatusOK && validationResponse.Data.Code == constant.SEERBIT_SUCCESS_CODE
	if pending {
		log.Println("payment successfully validated.")
	}

	return validationResponse, nil
}

func isInitialized(validation *Validation) (bool, any, error) {
	if !validation.Client.IsInitialized() {
		return true, nil, errors.New(constant.INITIALIZATION_ERROR)
	}
	return false, nil, nil
}
