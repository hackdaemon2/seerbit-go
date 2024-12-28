package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/account"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
)

func main() {
	processAccountPayment()
}

func processAccountPayment() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	account := account.NewAccount(seerBitClient)
	accountPayload := model.AccountPayload{
		PublicKey:          seerBitClient.PublicKey,
		Amount:             "100.00",
		Fee:                "10.00",
		FullName:           "John Doe",
		MobileNumber:       "08037456590",
		Currency:           "NGN",
		Country:            "NG",
		PaymentReference:   "paymentRef",
		Email:              "johndoe@gmail.com",
		ProductId:          "Foods",
		ProductDescription: "Uba Account Transaction",
		ClientAppCode:      "kpp64",
		ChannelType:        "BANK_ACCOUNT",
		RedirectUrl:        "https://checkout.seerbit.com",
		DeviceType:         "Apple Laptop",
		SourceIP:           "127.0.0.1:3456",
		AccountName:        "John S Doe",
		AccountNumber:      "1234567890",
		BankCode:           "033",
		BVN:                "12345678901",
		DateOfBirth:        "04011984",
		Retry:              "false",
		InvoiceNumber:      "1234567891abc123ac",
	}

	response, err := account.Pay(accountPayload)
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
