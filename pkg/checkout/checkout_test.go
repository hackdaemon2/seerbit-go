package checkout_test

import (
	"errors"
	"testing"

	"github.com/hackdaemon2/seerbit-go/pkg/account"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const API_BASE_URL = "https://api.seerbit.com"

// MockPaymentEngine is a mock implementation of PaymentEngine for testing.
type MockPaymentEngine struct {
	mock.Mock
}

func (m *MockPaymentEngine) ProcessPayment(
	payload any,
	paymentUrl, successCode string,
	authType client.AuthType) (any, error) {
	args := m.Called(payload, paymentUrl, successCode, authType)
	return args.Get(0), args.Error(1)
}

func TestNewAccount(t *testing.T) {
	mockClient := &client.SeerBitClient{BaseUrl: API_BASE_URL}
	accountInstance := account.NewAccount(mockClient)

	assert.NotNil(t, accountInstance)
	assert.Equal(t, mockClient, accountInstance.Client)
	assert.NotNil(t, accountInstance.PaymentEngine)
}

func TestPay_Success(t *testing.T) {
	mockClient := &client.SeerBitClient{BaseUrl: API_BASE_URL}
	mockPaymentEngine := new(MockPaymentEngine)

	accountInstance := &account.Account{
		Client:        mockClient,
		PaymentEngine: mockPaymentEngine,
	}

	payload := model.AccountPayload{
		Amount:             "1000",
		Currency:           "USD",
		ProductDescription: "Test payment",
	}

	paymentUrl := mockClient.BaseUrl + "/payments/initiates"
	mockPaymentEngine.On("ProcessPayment", payload, paymentUrl, client.SEERBIT_PENDING_CODE, client.Bearer).
		Return("success", nil)

	result, err := accountInstance.Pay(payload)

	assert.NoError(t, err)
	assert.Equal(t, "success", result)
	mockPaymentEngine.AssertExpectations(t)
}

func TestPay_InvalidPayload(t *testing.T) {
	mockClient := &client.SeerBitClient{BaseUrl: API_BASE_URL}
	mockPaymentEngine := new(MockPaymentEngine)

	accountInstance := &account.Account{
		Client:        mockClient,
		PaymentEngine: mockPaymentEngine,
	}

	invalidPayload := map[string]interface{}{
		"amount":   1000,
		"currency": "USD",
	}

	result, err := accountInstance.Pay(invalidPayload)

	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid payload for Account")
}

func TestPay_ProcessPaymentError(t *testing.T) {
	mockClient := &client.SeerBitClient{BaseUrl: API_BASE_URL}
	mockPaymentEngine := new(MockPaymentEngine)

	accountInstance := &account.Account{
		Client:        mockClient,
		PaymentEngine: mockPaymentEngine,
	}

	payload := model.AccountPayload{
		Amount:             "1000",
		Currency:           "USD",
		ProductDescription: "Test payment",
	}

	paymentUrl := mockClient.BaseUrl + "/payments/initiates"
	mockPaymentEngine.On("ProcessPayment", payload, paymentUrl, client.SEERBIT_PENDING_CODE, client.Bearer).
		Return(nil, errors.New("payment failed"))

	result, err := accountInstance.Pay(payload)

	assert.Nil(t, result)
	assert.EqualError(t, err, "payment failed")
	mockPaymentEngine.AssertExpectations(t)
}
