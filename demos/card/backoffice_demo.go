package main

import (
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/card"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
)

const ERROR_MESSAGE = "Error making payment: %v"

func RefundPayment(paymentReference, description string) {
	seerBitClient := getClient()

	card := card.NewCard(seerBitClient)
	refundPayload := formulatePayload(seerBitClient.PublicKey, paymentReference, description)

	response, err := card.Refund(refundPayload)
	if err != nil {
		log.Fatalf(ERROR_MESSAGE, err)
	}

	processResponse(response)
}

func CapturePayment(paymentReference, description string) {
	seerBitClient := getClient()

	card := card.NewCard(seerBitClient)
	refundPayload := formulatePayload(seerBitClient.PublicKey, paymentReference, description)

	response, err := card.Capture(refundPayload)
	if err != nil {
		log.Fatalf(ERROR_MESSAGE, err)
	}

	processResponse(response)
}

func CancelPayment(paymentReference, description string) {
	seerBitClient := getClient()

	card := card.NewCard(seerBitClient)
	refundPayload := formulatePayload(seerBitClient.PublicKey, paymentReference, description)

	response, err := card.Cancel(refundPayload)
	if err != nil {
		log.Fatalf(ERROR_MESSAGE, err)
	}

	processResponse(response)
}

func getClient() *client.SeerBitClient {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}
	return seerBitClient
}

func formulatePayload(publicKey, reference, productDescription string) model.PreauthorizationPayload {
	return model.PreauthorizationPayload{
		PublicKey:          publicKey,
		Amount:             "100",
		Currency:           "NGN",
		Country:            "NG",
		ProductDescription: productDescription,
		PaymentReference:   reference,
	}
}

func processResponse(response any) {
	switch resp := response.(type) {
	case model.PaymentResponse:
		printJson(resp)
	case model.ErrorResponse:
		log.Printf("Payment failed: %v", resp)
	default:
		log.Printf("Unexpected response type: %T", resp)
	}
}
