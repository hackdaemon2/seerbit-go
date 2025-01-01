package account_test

import (
	"testing"

	"github.com/hackdaemon2/seerbit-go/pkg/account"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/stub"
	"github.com/stretchr/testify/assert"
)

func TestNewVirtualAccount(t *testing.T) {
	mockClient := &client.SeerBitClient{BaseUrl: stub.API_BASE_URL}
	virtualAccount := account.NewVirtualAccount(mockClient)
	assert.NotNil(t, virtualAccount)
	assert.Equal(t, mockClient, virtualAccount.Client)
}

func TestCreateVirtualAccount_InvalidPayload(t *testing.T) {
	mockClient := &client.SeerBitClient{}
	virtualAccount := &account.VirtualAccount{Client: mockClient}

	invalidPayload := "invalid-payload"

	result, err := virtualAccount.Create(invalidPayload)

	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid payload for Virtual Account")
}

func TestCreateVirtualAccount_UninitializedClient(t *testing.T) {
	mockClient := &client.SeerBitClient{
		BaseUrl: stub.API_BASE_URL,
	}

	virtualAccount := &account.VirtualAccount{Client: mockClient}
	payload := model.VirtualAccountPayload{Currency: "USD"}

	result, err := virtualAccount.Create(payload)

	assert.Nil(t, result)
	assert.EqualError(t, err, client.INITIALIZATION_ERROR)
}

func TestGetPayments_InvalidAccountNumber(t *testing.T) {
	mockClient := &client.SeerBitClient{}
	virtualAccount := &account.VirtualAccount{Client: mockClient}

	result, err := virtualAccount.GetPayments("")

	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid account number for Virtual Account")
}

func TestRunWithReference_InvalidReference(t *testing.T) {
	mockClient := &client.SeerBitClient{}
	virtualAccount := &account.VirtualAccount{Client: mockClient}

	result, err := virtualAccount.GetVirtuaAccount("")

	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid payment reference for Virtual Account")
}
