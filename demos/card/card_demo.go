package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/card"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
)

func main() {
	processCardPayment()
}

func processCardPayment() {
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
		ProductDescription: "Uba Card Transaction",
		ClientAppCode:      "kpp64",
		ChannelType:        "MasterCard",
		PaymentType:        "CARD",
		RedirectUrl:        "https://checkout.seerbit.com",
		DeviceType:         "Apple Laptop",
		SourceIP:           "127.0.0.1:3456",
		Retry:              "false",
		InvoiceNumber:      "1234567891abc123ac",
		IsCvv:              "true",
		Pin:                "1234",
		CardNumber:         "5123450000000008",
		Cvv:                "100",
		ExpiryMonth:        "05",
		ExpiryYear:         "25",
	}

	response, err := card.Pay(cardPayload)
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

func printJson(response any) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}
	fmt.Println(string(jsonResponse))
}