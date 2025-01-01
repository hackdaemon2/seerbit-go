package momo

import (
	"errors"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/payment"
	"github.com/hackdaemon2/seerbit-go/pkg/validation"
)

// MobileMoney provides functionality to handle mobile money related operations,
// such as initiating payments using SeerBit's payment engine.
//
// For more details see https://seerbit.github.io/openapi/#operation/InitiatePayment
type MobileMoney struct {
	Client        *client.SeerBitClient
	PaymentEngine *payment.PaymentEngine
}

// NewMobileMoney creates a new MobileMoney instance.
// It initializes the MobileMoney with the provided SeerBitClient and a PaymentEngine.
//
// Parameters:
// - client: An instance of SeerBitClient to handle API interactions.
//
// Returns:
// - A pointer to the initialized MobileMoney struct.
func NewMobileMoney(client *client.SeerBitClient) *MobileMoney {
	return &MobileMoney{
		Client:        client,
		PaymentEngine: payment.NewPaymentEngine(client),
	}
}

// Pay processes a payment using the Mobile Money payment flow.
//
// Parameters:
// - payload: The payment data, expected to be of type model.MobileMoneyPayload.
//
// Returns:
// - (any): The response from the payment engine, or nil if an error occurs.
// - (error): An error if the payload is invalid or the payment processing fails.
//
// For more details see https://seerbit.github.io/openapi/#operation/InitiatePayment
func (momo *MobileMoney) Pay(payload any) (any, error) {
	mobileMoneyPayload, ok := payload.(model.MobileMoneyPayload)
	if !ok {
		return nil, errors.New("invalid payload for MOMO")
	}
	paymentUrl := momo.Client.BaseUrl + "/payments/initiates"
	return momo.PaymentEngine.ProcessPayment(mobileMoneyPayload, paymentUrl, client.SEERBIT_PENDING_CODE, client.Bearer)
}

// Finalize validates a payment using the Mobile Money payment flow.
//
// Parameters:
// - reference: The payment linking reference.
// - otp: one-time password from the Mobile Provider.
//
// Returns:
// - (any): The response from the payment engine, or nil if an error occurs.
// - (error): An error if the payload is invalid or the payment processing fails.
//
// For more details see https://seerbit.github.io/openapi/#operation/InitiatePayment
func (momo *MobileMoney) Finalize(reference, otp string) (any, error) {
	if reference == "" {
		return nil, errors.New("reference is required")
	}

	if otp == "" {
		return nil, errors.New("otp is required")
	}

	validate := validation.NewValidation(momo.Client)
	validationPayload := model.ValidationPayload{
		LinkingReference: reference,
		Otp:              otp,
	}

	return validate.Validate(validationPayload)
}
