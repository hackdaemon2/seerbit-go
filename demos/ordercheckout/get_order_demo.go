package main

import (
	"log"

	"github.com/hackdaemon2/seerbit-go/pkg/client"
	"github.com/hackdaemon2/seerbit-go/pkg/model"
	"github.com/hackdaemon2/seerbit-go/pkg/order"
)

const (
	REQUEST_FAILED_ERR        = "Requuest failed: %v"
	PUBLIC_KEY                = "your_public_key"
	PRIVATE_KEY               = "your_private_key"
	INITIALIZATION_ERROR      = "Failed to initialize SeerBit client: %v"
	UNEXPECTED_RESPONSE_ERROR = "Unexpected response type: %T"
)

func GetOrderPaymentRefDemo() {
	seerBitClient, err := client.NewSeerBitClient(PUBLIC_KEY, PRIVATE_KEY)
	if err != nil {
		log.Fatalf(INITIALIZATION_ERROR, err)
	}

	order := order.NewOrder(seerBitClient)

	response, err := order.GetOrdersByPaymentReference("payment-NH039920")
	if err != nil {
		log.Fatalf("Error retrieving order: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		if resp.Data.Code == client.SEERBIT_SUCCESS_CODE {
			log.Println("successful")
		} else {
			log.Printf(REQUEST_FAILED_ERR, resp)
		}
	case model.ErrorResponse:
		log.Printf(REQUEST_FAILED_ERR, resp)
	default:
		log.Printf(UNEXPECTED_RESPONSE_ERROR, resp)
	}

}

func GetOrdersDemo() {
	seerBitClient, err := client.NewSeerBitClient(PUBLIC_KEY, PRIVATE_KEY)
	if err != nil {
		log.Fatalf(INITIALIZATION_ERROR, err)
	}

	order := order.NewOrder(seerBitClient)

	response, err := order.GetOrders()
	if err != nil {
		log.Fatalf("Error creating order: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		if resp.Data.Code == client.SEERBIT_SUCCESS_CODE {
			log.Println("successful")
		} else {
			log.Printf(REQUEST_FAILED_ERR, resp)
		}
	case model.ErrorResponse:
		log.Printf(REQUEST_FAILED_ERR, resp)
	default:
		log.Printf(UNEXPECTED_RESPONSE_ERROR, resp)
	}

}

func GetOrderOrderIdDemo() {
	seerBitClient, err := client.NewSeerBitClient(PUBLIC_KEY, PRIVATE_KEY)
	if err != nil {
		log.Fatalf(INITIALIZATION_ERROR, err)
	}

	order := order.NewOrder(seerBitClient)

	response, err := order.GetOrdersByOrderId("order-NH039920")
	if err != nil {
		log.Fatalf("Error retrieving order: %v", err)
	}

	switch resp := response.(type) {
	case model.PaymentResponse:
		if resp.Data.Code == client.SEERBIT_SUCCESS_CODE {
			log.Println("successful")
		} else {
			log.Printf(REQUEST_FAILED_ERR, resp)
		}
	case model.ErrorResponse:
		log.Printf(REQUEST_FAILED_ERR, resp)
	default:
		log.Printf(UNEXPECTED_RESPONSE_ERROR, resp)
	}

}
