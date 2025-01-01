package main

import (
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/order"
)

func ProcessCreateOrderPostPaymentDemo() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	order := order.NewOrder(seerBitClient)
	orderPayload := createNewOrder(seerBitClient.PublicKey)

	response, err := order.PostCreate(orderPayload)
	if err != nil {
		log.Fatalf("Error creating order: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		if resp.Data.Code == client.SEERBIT_SUCCESS_CODE {
			log.Println("successful")
		} else {
			log.Printf("Request failed: %v", resp)
		}
	case model.ErrorResponse:
		log.Printf("Request failed: %v", resp)
	default:
		log.Printf("Unexpected response type: %T", resp)
	}
}

func createNewOrder(publicKey string) model.OrderPostPaymentPayload {
	return model.OrderPostPaymentPayload{
		PublicKey:        publicKey,
		PaymentReference: "reference-101",
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
}
