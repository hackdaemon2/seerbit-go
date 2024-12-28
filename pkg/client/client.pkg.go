package client

type authPayload struct {
	Key string `json:"key"`
}

type authData struct {
	Code            string          `json:"code"`
	Message         string          `json:"message"`
	EncryptedSecKey encryptedSecKey `json:"encryptedSecKey"`
}

type encryptedSecKey struct {
	EncryptedKey string `json:"encryptedKey"`
}

type authResponse struct {
	Status string   `json:"status"`
	Data   authData `json:"data"`
}

type authErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
