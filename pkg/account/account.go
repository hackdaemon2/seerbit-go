package account

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/constant"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/payment"
	"github.com/hackdaemon2/seerbit-go/util"
)

type Account struct {
	Client        *client.SeerBitClient
	paymentEngine *payment.PaymentEngine
}

func NewAccount(client *client.SeerBitClient) *Account {
	return &Account{
		Client:        client,
		paymentEngine: payment.NewPaymentEngine(client),
	}
}

func (account *Account) Pay(payload any) (any, error) {
	accountPayload, ok := payload.(model.AccountPayload)
	if !ok {
		return nil, errors.New("invalid payload for Account")
	}
	paymentUrl := account.Client.BaseUrl + "/payments/initiates"
	return account.paymentEngine.ProcessPayment(accountPayload, paymentUrl, constant.SEERBIT_PENDING_CODE, constant.Bearer)
}

func (account *Account) GetBanks() (any, error) {
	var getBanksResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	getBanksUrl := strings.Join([]string{account.Client.BaseUrl, "/banks/merchant/", account.Client.PublicKey}, "")
	httpRequest := util.HttpRequestData{
		Response:       getBanksResponse,
		ErrorResponse:  errorResponse,
		Url:            getBanksUrl,
		PublicKey:      account.Client.PublicKey,
		PrivateKey:     account.Client.PrivateKey,
		AuthType:       constant.Bearer,
		Authentication: account.Client.BearerToken,
	}

	resp, err := httpRequest.HttpGet()
	if err != nil {
		return nil, fmt.Errorf(constant.ERROR_MESSAGE, err)
	}

	shouldReturn, checkoutError, err := httpRequest.IsErrorResponse(resp, errorResponse)
	if shouldReturn {
		return checkoutError, err
	}

	log.Println("bank list successfully retrieved.")
	return getBanksResponse, nil
}
