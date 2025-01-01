package recurring

import (
	"fmt"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/payment"
	"github.com/hackdaemon2/seerbit-go/util"
)

type Subscription struct {
	Client        *client.SeerBitClient
	PaymentEngine *payment.PaymentEngine
}

func NewSubscription(client *client.SeerBitClient) *Subscription {
	return &Subscription{
		Client:        client,
		PaymentEngine: payment.NewPaymentEngine(client),
	}
}

func (subscription *Subscription) RecurringSubscription(payload model.CardPayload) (any, error) {
	paymentUrl := subscription.Client.BaseUrl + "/recurring/subscribes"
	return subscription.PaymentEngine.ProcessPayment(payload, paymentUrl, client.SEERBIT_PENDING_CODE, client.Bearer)
}

func (subscription *Subscription) GetSubscription(billingId string) (any, error) {
	var subscriptionResponse model.PaymentResponse
	var errorResponse model.ErrorResponse

	subscriptionUrl := fmt.Sprintf("%s/recurring/billingId/%s", subscription.Client.BaseUrl, billingId)
	httpRequest := util.HttpRequestData{
		Response:       subscriptionResponse,
		ErrorResponse:  errorResponse,
		Url:            subscriptionUrl,
		PublicKey:      subscription.Client.PublicKey,
		PrivateKey:     subscription.Client.PrivateKey,
		AuthType:       string(client.Bearer),
		Authentication: subscription.Client.BearerToken,
	}

	resp, err := httpRequest.HttpGet()
	if err != nil {
		return nil, fmt.Errorf(client.ERROR_MESSAGE, err)
	}

	shouldReturn, _, err := httpRequest.IsErrorResponse(resp, &errorResponse, &subscriptionResponse)
	if shouldReturn && len(errorResponse.Error) != 0 {
		return errorResponse, err
	}

	return subscriptionResponse, nil
}
