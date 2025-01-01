package main

import (
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/order"
)

func ProcessCreateOrderPrePaymentDemo() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	order := order.NewOrder(seerBitClient)
	orderPayload := model.OrderPaymentPayload{
		Email:            "john.doe@gmail.com",
		PublicKey:        seerBitClient.PublicKey,
		PaymentReference: "reference-101",
		FullName:         "John Doe",
		OrderType:        "BULK_BULK",
		MobileNumber:     "08030000002",
		CallbackURL:      "http://mvaa.ebs-rcm.com/payment",
		Country:          "NG",
		Currency:         "NGN",
		Amount:           "100",
		Orders: []model.OrderData{
			{
				OrderID:            "LA771",
				Currency:           "NGN",
				Amount:             "70",
				ProductID:          "939920KKL",
				ProductDescription: "FRSC Plate Number",
			},
			{
				OrderID:            "LA772",
				Currency:           "NGN",
				Amount:             "30",
				ProductID:          "20BHJ9292",
				ProductDescription: "Android Wristwatch",
			},
		},
	}

	response, err := order.Create(orderPayload)
	if err != nil {
		log.Fatalf("Error creating order: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		if resp.Data.Code == client.SEERBIT_PENDING_CODE || resp.Data.Code == client.SEERBIT_SUCCESS_CODE {
			log.Printf("redirect link => %s", resp.Data.Payment.RedirectLink)
		} else {
			log.Printf("Payment failed: %v", resp)
		}
	case model.ErrorResponse:
		log.Printf("Payment failed: %v", resp)
	default:
		log.Printf("Unexpected response type: %T", resp)
	}

}
