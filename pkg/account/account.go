package account

import (
	"errors"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/payment"
)

// Account provides functionality to handle account-related operations,
// such as initiating payments using SeerBit's payment engine.
//
// For more details see https://seerbit.github.io/openapi/#operation/InitiatePayment
type Account struct {
	Client        *client.SeerBitClient          // Client instance for interacting with SeerBit APIs.
	PaymentEngine payment.PaymentEngineProcessor // Payment engine for processing payments.
}

// NewAccount creates a new Account instance.
// It initializes the Account with the provided SeerBitClient and a PaymentEngine.
//
// Parameters:
// - client (*client.SeerBitClient): An instance of SeerBitClient to handle API interactions.
//
// Returns:
// - (*Account): A pointer to the initialized Account struct.
func NewAccount(client *client.SeerBitClient) *Account {
	return &Account{
		Client:        client,
		PaymentEngine: payment.NewPaymentEngine(client),
	}
}

// Pay processes a payment using the Account payment flow.
//
// Parameters:
// - payload: The payment data, expected to be of type model.AccountPayload.
//
// Returns:
// - (any): The response from the payment engine, or nil if an error occurs.
// - (error): An error if the payload is invalid or the payment processing fails.
//
// For more details see https://seerbit.github.io/openapi/#operation/InitiatePayment
func (account *Account) Pay(payload any) (any, error) {
	// Validate that the payload is of type model.AccountPayload.
	accountPayload, ok := payload.(model.AccountPayload)
	if !ok {
		return nil, errors.New("invalid payload for Account")
	}

	// Construct the payment URL using the SeerBit client's base URL.
	paymentUrl := account.Client.BaseUrl + "/payments/initiates"

	// Process the payment using the PaymentEngine.
	// The SeerBit pending code and authentication type (Bearer) are passed along.
	return account.PaymentEngine.ProcessPayment(accountPayload, paymentUrl, client.SEERBIT_PENDING_CODE, client.Bearer)
}
