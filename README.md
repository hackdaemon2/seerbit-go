# seerbit-go - SeerBit Payment API Library for GOLANG

## Features

The Library supports all APIs under the following services:

* Payments via API (mobile money, cards, account, etc.)
* Recurring Payments
* Transaction Status

## Getting Started

A full getting started guide for integrating SeerBit can be found at [getting started docs](https://doc.seerbit.com).

## Documentation

The documentation, installation guide, detailed description of the SeerBit API and all of its features is [available on the documentation website](https://doc.seerbit.com/api/library)

## Requirements

* GO 1.21.9

## Installation

Run this command on the terminal:

```code
go get https://github.com/hackdaemon2/seerbit-go
```

## Contributing

You can contribute to this repository so that anyone can benefit from it:

* Improved features
* Resolved bug fixes and issues

## Demos  

You can also check the [demos](https://github.com/hackdaemon2/seerbit-go/tree/master/demos)

## Using the Library

### Initiate Standard Checkout Payment

Instantiate a client and set the parameters.

```go
  package main

  import (
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/hackdaemon2/seerbit-go/pkg/checkout"
    "github.com/hackdaemon2/seerbit-go/pkg/client"
    "github.com/hackdaemon2/seerbit-go/pkg/model"
  )

  func init() {
    if err := godotenv.Load(".env"); err != nil {
      log.Fatal("Error loading .env file")
    }
  }

  func main() {
    processSeerBitPayment()
  }

  func processSeerBitPayment() {
    seerBitClient, err := client.NewSeerBitClient(os.Getenv("SEERBIT_PUBLIC_KEY"), os.Getenv("SEERBIT_PRIVATE_KEY"))
    if err != nil {
      log.Fatalf("Failed to initialize SeerBit client: %v", err)
    }

    standardCheckout := checkout.NewCheckout(seerBitClient)
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

    response, err := standardCheckout.Pay(cardPayload)
    if err != nil {
      log.Fatalf("Error making payment: %v", err)
    }

    switch resp := response.(type) {
    case model.PaymentResponse:
      if resp.Data.Code == client.SEERBIT_PENDING_CODE || resp.Data.Code == client.SEERBIT_SUCCESS_CODE {
        log.Printf("Payment pending: %v", resp)
      } else {
        log.Printf("Payment failed: %v", resp)
      }
    case model.ErrorResponse:
      log.Printf("Payment failed: %v", resp)
    default:
      log.Printf("Unexpected response type: %T", resp)
    }
  }

```

Find more examples [here](https://github.com/hackdaemon2/seerbit-go/tree/master/demos).

## Licence

The MIT License (MIT). For more information, see the LICENSE file.
