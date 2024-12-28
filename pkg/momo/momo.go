package momo

import (
	"errors"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/constant"
	"github.com/hackdaemon2/seerbit-go/pkg/payment"
)

type MobileMoney struct {
	Client        *client.SeerBitClient
	paymentEngine *payment.PaymentEngine
}

func NewMobileMoney(client *client.SeerBitClient) *MobileMoney {
	return &MobileMoney{
		Client:        client,
		paymentEngine: payment.NewPaymentEngine(client),
	}
}

func (momo *MobileMoney) Pay(payload any) (any, error) {
	mobileMoneyPayload, ok := payload.(MobileMoneyPayload)
	if !ok {
		return nil, errors.New("invalid payload for MOMO")
	}
	paymentUrl := momo.Client.BaseUrl + "/payments/initiates"
	return momo.paymentEngine.ProcessPayment(mobileMoneyPayload, paymentUrl, constant.SEERBIT_PENDING_CODE, constant.Bearer)
}
