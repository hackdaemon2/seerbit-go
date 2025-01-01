package util

import (
	"github.com/go-resty/resty/v2"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
)

type PaymentService interface {
	Pay(any) (any, error)
}

type ValidationService interface {
	Validate(model.ValidationPayload) (any, error)
	Verify(reference string) (any, error)
}

type FinalizeService interface {
	Finalize(reference, otp string) (any, error)
}

type PreauthorizationService interface {
	DirectCharge(model.CardPayload) (any, error)
	Authorize(model.CardPayload) (any, error)
	Capture(model.PreauthorizationPayload) (any, error)
	Refund(model.PreauthorizationPayload) (any, error)
	Cancel(model.PreauthorizationPayload) (any, error)
}

type HttpService interface {
	HttpGet() (*resty.Response, error)
	HttpPost() (*resty.Response, error)
	HttpPut() (*resty.Response, error)
	IsErrorResponse(*resty.Response, *model.ErrorResponse, *model.PaymentResponse) (bool, any, error)
}
