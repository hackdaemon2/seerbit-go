package main

import (
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/recurring"
)

func GetSubscriptionDemo() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	subscription := recurring.NewSubscription(seerBitClient)

	response, err := subscription.GetSubscription("billing-id-REF")
	if err != nil {
		log.Fatalf("Error making payment: %v", err)
	}

	processResponse(response)
}

func ProcessRecurringSubscription() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	subscription := recurring.NewSubscription(seerBitClient)
	cardPayload := model.CardPayload{
		StartDate:          "2020-02-25 00:00:00",
		PlanId:             "3484839",
		CustomerId:         "4432KKe3",
		BillingCycle:       "WEEKLY",
		BillingPeriod:      "1",
		SubscriptionAmount: true,
		PublicKey:          seerBitClient.PublicKey,
		Amount:             "100.00",
		MobileNumber:       "08037456590",
		Currency:           "NGN",
		Country:            "NG",
		PaymentReference:   "paymentRef",
		Email:              "johndoe@gmail.com",
		ProductId:          "Foods",
		CallbackUrl:        "www.checkout.com/redirect",
		ProductDescription: "Uba Card Transaction",
		CardName:           "John Doe",
		Type:               "3DSECURE",
		Pin:                "1234",
		CardNumber:         "5123450000000008",
		Cvv:                "100",
		ExpiryMonth:        "05",
		ExpiryYear:         "25",
	}

	response, err := subscription.RecurringSubscription(cardPayload)
	if err != nil {
		log.Fatalf("Error making payment: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		printJson(resp)
		if len(resp.Data.Payment.RedirectUrl) != 0 {
			redirectUrl := resp.Data.Payment.RedirectUrl
			log.Printf("redirect link => %s", redirectUrl)
		}
	case model.ErrorResponse:
		log.Printf("Payment failed: %v", resp)
	default:
		log.Printf("Unexpected response type: %T", resp)
	}
}
