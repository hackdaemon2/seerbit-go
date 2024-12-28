package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/hackdaemon2/seerbit-go/pkg/constant"
	"github.com/hackdaemon2/seerbit-go/util"
)

type SeerBitClient struct {
	PublicKey     string
	PrivateKey    string
	BearerToken   string
	isInitialised bool
	BaseUrl       string
}

// NewSeerBitClient creates a new Client instance with required fields
func NewSeerBitClient(publicKey, privateKey string) (*SeerBitClient, error) {
	if publicKey == "" {
		return nil, errors.New("public key must be set")
	}

	if privateKey == "" {
		return nil, errors.New("private key must be set")
	}

	client := &SeerBitClient{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		BaseUrl:    constant.SEERBIT_BASE_URL,
	}

	err := client.initialize()
	return client, err
}

// initialize initializes the client and sets the BearerToken
func (client *SeerBitClient) initialize() error {
	if client.isInitialised {
		log.Println("Client is already initialized.")
		return nil
	}

	var authResp authResponse
	var authErrResp authErrorResponse

	url := client.BaseUrl + "/encrypt/keys"
	authPayload := &authPayload{Key: fmt.Sprintf("%s.%s", client.PrivateKey, client.PublicKey)}

	httpRequest := util.HttpRequestData{
		PrivateKey:    client.PrivateKey,
		PublicKey:     client.PublicKey,
		Payload:       authPayload,
		Response:      authResp,
		ErrorResponse: authErrResp,
		Url:           url,
		AuthType:      constant.NoAuth,
	}

	resp, err := httpRequest.HttpPost()
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	if resp.StatusCode() == http.StatusOK && authResp.Data.Code == constant.SEERBIT_SUCCESS_CODE {
		client.BearerToken = authResp.Data.EncryptedSecKey.EncryptedKey
		client.isInitialised = true
		log.Println("Client successfully initialized.")
		return nil
	}

	jsonResponse, _ := json.Marshal(authErrResp)
	return fmt.Errorf("authentication failed: %s", jsonResponse)
}

// IsInitialized returns whether the client is initialized
func (client *SeerBitClient) IsInitialized() bool {
	return client.isInitialised
}
