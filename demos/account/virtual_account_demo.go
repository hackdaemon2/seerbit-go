package main

import (
	"fmt"
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/account"
	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
)

func CreateVirtualAccount() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	virtualAccount := account.NewVirtualAccount(seerBitClient)
	accountPayload := model.VirtualAccountPayload{
		PublicKey:              seerBitClient.PublicKey,
		FullName:               "John Doe",
		Currency:               "NGN",
		Country:                "NG",
		Email:                  "johndoe@gmail.com",
		BankVerificationNumber: "100299010",
		Reference:              "ref_200010",
	}

	response, err := virtualAccount.Create(accountPayload)
	if err != nil {
		log.Fatalf("Error making payment: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		fmt.Println("account number => ", resp.Data.Payment.AccountNumber)
		fmt.Println("wallet name => ", resp.Data.Payment.WalletName)
	case model.ErrorResponse:
		log.Printf("Payment failed: %v", resp)
	default:
		log.Printf("Unexpected response type: %T", resp)
	}
}

func GetPayments() {
	seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
	if err != nil {
		log.Fatalf("Failed to initialize SeerBit client: %v", err)
	}

	virtualAccount := account.NewVirtualAccount(seerBitClient)

	response, err := virtualAccount.GetPayments("ref_200010")
	if err != nil {
		log.Fatalf("Error making payment: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		for _, payment := range resp.Data.Payload {
			fmt.Println(payment.CreditAccountNumber)
		}
	case model.ErrorResponse:
		log.Printf("Payment failed: %v", resp)
	default:
		log.Printf("Unexpected response type: %T", resp)
	}

}
