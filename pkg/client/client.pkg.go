package client

const (
	SEERBIT_SUCCESS_CODE          = "00"
	SEERBIT_PENDING_CODE          = "S20"
	SEERBIT_BASE_URL              = "https://seerbitapi.com/api/v2"
	INVALID_PAYLOAD               = "invalid payload for "
	ERROR_MESSAGE                 = "error making request: %w"
	SEERBIT_HASH_TYPE             = "sha256"
	INITIALIZATION_ERROR          = "client has not been initialized"
	NoAuth               AuthType = "NoAuth"
	Bearer               AuthType = "Bearer"
	Basic                AuthType = "Basic"
)

type AuthType string

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
