package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/validation"
)

func ProcessVerify() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	validation := validation.NewValidation(seerBitClient)
	response, err := validation.Verify("pay-2020300-WQJ")
	if err != nil {
		log.Fatalf("Error making payment: %v", err)
	}

	processResponse(response)
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

func printJson(response any) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}
	fmt.Println(string(jsonResponse))
}
