package util

import "github.com/hackdaemon2/seerbit-go/pkg/model"

type PaymentService interface {
	Pay(any) (any, error)
}

type ValidationService interface {
	Validate(model.ValidationPayload) (any, error)
	Verify(reference string) (any, error)
}

type PreauthorizationService interface {
	DirectCharge(model.CardPayload) (any, error)
	Authorize(model.CardPayload) (any, error)
	Capture(model.PreauthorizationPayload) (any, error)
	Refund(model.PreauthorizationPayload) (any, error)
	Cancel(model.PreauthorizationPayload) (any, error)
}
