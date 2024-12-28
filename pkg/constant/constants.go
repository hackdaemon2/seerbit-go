package constant

type AuthType string

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
