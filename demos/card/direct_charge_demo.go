package main

import (
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/card"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
)

func ProcessDirectChargePayment() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	card := card.NewCard(seerBitClient)
	cardPayload := model.CardPayload{
		PublicKey:          seerBitClient.PublicKey,
		Amount:             "100.00",
		FullName:           "John Doe",
		MobileNumber:       "08037456590",
		Currency:           "NGN",
		Country:            "NG",
		PaymentReference:   "paymentRef",
		Email:              "johndoe@gmail.com",
		ProductId:          "Foods",
		ProductDescription: "Uba Account Transaction",
		CardNumber:         "5123450000000008",
		Cvv:                "100",
		ExpiryMonth:        "05",
		ExpiryYear:         "25",
	}

	response, err := card.DirectCharge(cardPayload)
	if err != nil {
		log.Fatalf("Error making payment: %v", err)
	}

	processResponse(response)
}
