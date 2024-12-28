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

### Initiate Account Option

Instantiate a client and set the parameters.

```go
  package main

  import (
      "encoding/json"
      "fmt"
      "log"

      "github.com/hackdaemon2/seerbit-go/account"
      "github.com/hackdaemon2/seerbit-go/client"
      "github.com/hackdaemon2/seerbit-go/model"
  )

  func main() {
      seerBitClient, err := client.NewSeerBitClient("your_public_key", "your_private_key")
      if err != nil {
          log.Fatalf("Failed to initialize SeerBit client: %v", err)
      }

      accountPayload := &account.AccountPayload{
          PublicKey:         "your_public_key",
          Amount:            "100.00",
          Fee:               "10.00",
          FullName:          "John Doe",
          MobileNumber:      "08037456590",
          Currency:          "NGN",
          Country:           "NG",
          PaymentReference:  "paymentRef",
          Email:             "johndoe@gmail.com",
          ProductID:         "Foods",
          ProductDescription: "Uba Account Transaction",
          ClientAppCode:     "kpp64",
          ChannelType:       "BANK_ACCOUNT",
          RedirectUrl:       "https://checkout.seerbit.com",
          DeviceType:        "Apple Laptop",
          SourceIP:          "127.0.0.1:3456",
          AccountName:       "John S Doe",
          AccountNumber:     "1234567890",
          BankCode:          "033",
          BVN:               "12345678901",
          DateOfBirth:       "04011984",
          Retry:             "false",
          InvoiceNumber:     "1234567891abc123ac",
      }

      accountInstance := account.NewAccount(seerBitClient)

      response, err := accountInstance.Pay(accountPayload)
      if err != nil {
          log.Fatalf("Error making payment: %v", err)
      }

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
```

Find more examples [here](https://github.com/hackdaemon2/seerbit-go/tree/master/demos).

## Licence

GNU General Public License. For more information, see the LICENSE file.
