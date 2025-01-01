package stub

import (
	"github.com/go-resty/resty/v2"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/stretchr/testify/mock"
)

const API_BASE_URL = "https://api.seerbit.com"
const PRIVATE_KEY = "test-private-key"
const PUBLIC_KEY = "test-public-key"
const BEARER_TOKEN = "test-bearer-token"

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

// MockHttpRequest simulates HTTP requests for testing.
type MockHttpRequest struct {
	mock.Mock
}

func (m *MockHttpRequest) HttpGet() (*resty.Response, error) {
	return returnMockResponse(m)
}

func (m *MockHttpRequest) HttpPut() (*resty.Response, error) {
	return returnMockResponse(m)
}

func (m *MockHttpRequest) HttpPost() (*resty.Response, error) {
	return returnMockResponse(m)
}

func (m *MockHttpRequest) IsErrorResponse(
	resp *resty.Response,
	errResp *model.ErrorResponse,
	payResp *model.PaymentResponse) (bool, any, error) {
	args := m.Called(resp, errResp, payResp)
	return false, args.Get(2).(*model.PaymentResponse), nil
}

func returnMockResponse(m *MockHttpRequest) (*resty.Response, error) {
	args := m.Called()
	return args.Get(0).(*resty.Response), args.Error(1)
}
