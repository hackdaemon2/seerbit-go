package main

import (
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/constant"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/momo"
	"github.com/hackdaemon2/seerbit-go/pkg/validation"
)

func ProcessMomoPayment() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	mobileMoney := momo.NewMobileMoney(seerBitClient)
	momoPayload := model.MobileMoneyPayload{
		PublicKey:          seerBitClient.PublicKey,
		Amount:             "100.00",
		FullName:           "John Doe",
		MobileNumber:       "08037456590",
		Currency:           "NGN",
		DeviceType:         "my_device_type",
		Network:            "MTN",
		VoucherCode:        "V8281E",
		Fee:                "1",
		SourceIP:           "127.0.0.1",
		Country:            "NG",
		PaymentReference:   "paymentRef",
		Email:              "johndoe@gmail.com",
		ProductId:          "Foods",
		ProductDescription: "OPay Momo Transaction",
	}

	response, err := mobileMoney.Pay(momoPayload)
	if err != nil {
		log.Fatalf("Error making payment: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		if resp.Data.Code == constant.SEERBIT_PENDING_CODE {
			log.Printf("linking reference => %s", resp.Data.Payment.LinkingReference)
		}
	case model.ErrorResponse:
		log.Printf("Payment failed: %v", resp)
	default:
		log.Printf("Unexpected response type: %T", resp)
	}
}

func ProcessValidateMomo() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	validation := validation.NewValidation(seerBitClient)
	validatePayload := model.ValidationPayload{
		LinkingReference: "32291919",
		Otp:              "2991",
	}

	response, err := validation.Validate(validatePayload)
	if err != nil {
		log.Fatalf("Error making payment: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		if resp.Data.Code == constant.SEERBIT_SUCCESS_CODE {
			log.Printf("payment successful => %s", resp.Data.Code)
		}
	case model.ErrorResponse:
		log.Printf("Payment failed: %v", resp)
	default:
		log.Printf("Unexpected response type: %T", resp)
	}
}
