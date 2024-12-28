package main

import (
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/checkout"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/constant"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
)

func ProcessStandardCheckoutPayment() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	checkout := checkout.NewCheckout(seerBitClient)
	cardPayload := model.CheckoutPayload{
		PublicKey:          seerBitClient.PublicKey,
		Amount:             "100.00",
		Currency:           "NGN",
		Country:            "NG",
		PaymentReference:   "paymentRef",
		Email:              "johndoe@gmail.com",
		ProductId:          "Foods",
		ProductDescription: "Checkout Payment Transaction",
	}

	response, err := checkout.Pay(cardPayload)
	if err != nil {
		log.Fatalf("Error making payment: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		if resp.Data.Code == constant.SEERBIT_PENDING_CODE || resp.Data.Code == constant.SEERBIT_SUCCESS_CODE {
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