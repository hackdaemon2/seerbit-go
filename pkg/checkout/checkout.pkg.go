package standardcheckout

type CheckoutHashPayload struct {
	CheckoutPayload
	Hash     string `json:"hash,omitempty"`
	HashType string `json:"hashType,omitempty"`
}

type CheckoutPayload struct {
	PublicKey          string `json:"publicKey,omitempty"`
	Amount             string `json:"amount,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Country            string `json:"country,omitempty"`
	PaymentReference   string `json:"paymentReference,omitempty"`
	Email              string `json:"email,omitempty"`
	ProductId          string `json:"productId,omitempty"`
	ProductDescription string `json:"productDescription,omitempty"`
	CallbackUrl        string `json:"callbackUrl,omitempty"`
}
