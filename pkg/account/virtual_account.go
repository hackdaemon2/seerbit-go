package account

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/util"
)

// VirtualAccount provides functionality for managing virtual accounts
// and interacting with the SeerBit API for creating, retrieving, and deleting virtual accounts.
//
// For more details see https://doc.seerbit.com/online-payment/payment-features/virtual-accounts
type VirtualAccount struct {
	Client *client.SeerBitClient // Client instance to interact with SeerBit APIs.
}

// NewVirtualAccount initializes a VirtualAccount instance.
//
// Parameters:
// - client (*client.SeerBitClient): An instance of SeerBitClient to handle API interactions.
//
// Returns:
// - (*VirtualAccount): A pointer to the initialized VirtualAccount struct.
//
// For more details see https://doc.seerbit.com/online-payment/payment-features/virtual-accounts
func NewVirtualAccount(client *client.SeerBitClient) *VirtualAccount {
	return &VirtualAccount{
		Client: client,
	}
}

// Create initializes a new virtual account using the provided payload.
//
// Parameters:
// - payload: The data for creating a virtual account, expected to be of type model.VirtualAccountPayload.
//
// Returns:
// - (any): The response from the SeerBit API, or nil if an error occurs.
// - (error): An error if the payload is invalid or the request fails.
//
// For more details see https://doc.seerbit.com/online-payment/payment-features/virtual-accounts
func (virtualAccount *VirtualAccount) Create(payload any) (any, error) {
	virtualAccountPayload, ok := payload.(model.VirtualAccountPayload)
	if !ok {
		return nil, errors.New("invalid payload for Virtual Account")
	}

	if !virtualAccount.Client.IsInitialized() {
		return nil, errors.New(client.INITIALIZATION_ERROR)
	}

	url := virtualAccount.Client.BaseUrl + "/virtual-accounts"
	return executeRequest(virtualAccountPayload, url, http.MethodPost, virtualAccount.Client)
}

// GetPayments retrieves payment information for a specific virtual account.
//
// Parameters:
// - accountNumber (string): The account number for which payments are to be retrieved.
//
// Returns:
// - (any): The response from the SeerBit API, or nil if an error occurs.
// - (error): An error if the account number is invalid or the request fails.
//
// For more details see https://doc.seerbit.com/online-payment/payment-features/virtual-accounts
func (virtualAccount *VirtualAccount) GetPayments(accountNumber string) (any, error) {
	if accountNumber == "" {
		return nil, errors.New("invalid account number for Virtual Account")
	}

	if !virtualAccount.Client.IsInitialized() {
		return nil, errors.New(client.INITIALIZATION_ERROR)
	}

	url := strings.Join([]string{
		virtualAccount.Client.BaseUrl,
		"/virtual-accounts/",
		virtualAccount.Client.PublicKey,
		"/",
		accountNumber,
	}, "")

	return executeRequest(nil, url, http.MethodGet, virtualAccount.Client)
}

// GetVirtuaAccount retrieves details of a virtual account using a payment reference.
//
// Parameters:
// - paymentReference (string): The reference associated with the virtual account.
//
// Returns:
// - (any): The response from the SeerBit API, or nil if an error occurs.
// - (error): An error if the reference is invalid or the request fails.
//
// For more details see https://doc.seerbit.com/online-payment/payment-features/virtual-accounts
func (virtualAccount *VirtualAccount) GetVirtuaAccount(paymentReference string) (any, error) {
	return virtualAccount.runWithReference(paymentReference, http.MethodGet)
}

// DeleteVirtuaAccount deletes a virtual account using a payment reference.
//
// Parameters:
// - paymentReference: The reference associated with the virtual account to be deleted.
//
// Returns:
// - (any): The response from the SeerBit API, or nil if an error occurs.
// - (error): An error if the reference is invalid or the request fails.
//
// For more details see https://doc.seerbit.com/online-payment/payment-features/virtual-accounts
func (virtualAccount *VirtualAccount) DeleteVirtuaAccount(paymentReference string) (any, error) {
	return virtualAccount.runWithReference(paymentReference, http.MethodDelete)
}

// runWithReference executes an API request using a payment reference and HTTP method.
//
// Parameters:
// - reference: The payment reference.
// - method: The HTTP method (GET, DELETE).
//
// Returns:
// - (any): The response from the SeerBit API, or nil if an error occurs.
// - (error): An error if the reference is invalid or the request fails.
func (virtualAccount *VirtualAccount) runWithReference(reference, method string) (any, error) {
	if reference == "" {
		return nil, errors.New("invalid payment reference for Virtual Account")
	}

	if !virtualAccount.Client.IsInitialized() {
		return nil, errors.New(client.INITIALIZATION_ERROR)
	}

	url := strings.Join([]string{virtualAccount.Client.BaseUrl, "/virtual-accounts/", reference}, "")
	return executeRequest(nil, url, method, virtualAccount.Client)
}

// executeRequest handles the execution of HTTP requests using Resty.
//
// Parameters:
// - payload: The request payload.
// - url: The API endpoint.
// - method: The HTTP method (POST, GET, DELETE).
// - seerBitClient: The SeerBitClient instance for authentication.
//
// Returns:
// - (any): The response from the SeerBit API, or nil if an error occurs.
// - (error): An error if the request or response processing fails.
func executeRequest(payload any, url, method string, seerBitClient *client.SeerBitClient) (any, error) {
	var virtualAccountResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	httpRequest := util.HttpRequestData{
		PrivateKey:     seerBitClient.PrivateKey,
		PublicKey:      seerBitClient.PublicKey,
		Payload:        payload,
		Response:       &virtualAccountResponse,
		ErrorResponse:  &errorResponse,
		Url:            url,
		Authentication: seerBitClient.BearerToken,
		AuthType:       string(client.Bearer),
	}

	var resp *resty.Response
	var err error

	switch method {
	case http.MethodPost:
		resp, err = httpRequest.HttpPost()
	case http.MethodGet:
		resp, err = httpRequest.HttpGet()
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	shouldReturn, _, err := httpRequest.IsErrorResponse(resp, &errorResponse, &virtualAccountResponse)
	if shouldReturn && len(errorResponse.Error) != 0 {
		return errorResponse, err
	}

	return virtualAccountResponse, nil
}
