package main

import (
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/validation"
)

func ProcessValidate() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	validation := validation.NewValidation(seerBitClient)
	response, err := validation.Validate(model.ValidationPayload{Otp: "392928", LinkingReference: "pay-2020300-WQJ"})
	if err != nil {
		log.Fatalf("Error making payment: %v", err)
	}

	processResponse(response)
}
