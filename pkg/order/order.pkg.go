package order

type OrderData struct {
	OrderID            string `json:"orderId,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Amount             string `json:"amount,omitempty"`
	ProductID          string `json:"productId,omitempty"`
	ProductDescription string `json:"productDescription,omitempty"`
}

type OrderPaymentPayload struct {
	Email            string  `json:"email,omitempty"`
	PublicKey        string  `json:"publicKey,omitempty"`
	PaymentReference string  `json:"paymentReference,omitempty"`
	FullName         string  `json:"fullName,omitempty"`
	OrderType        string  `json:"orderType,omitempty"`
	MobileNumber     string  `json:"mobileNumber,omitempty"`
	CallbackURL      string  `json:"callbackUrl,omitempty"`
	Country          string  `json:"country,omitempty"`
	Currency         string  `json:"currency,omitempty"`
	Amount           string  `json:"amount,omitempty"`
	Orders           []Order `json:"orders,omitempty"`
}

type OrderPostPaymentPayload struct {
	PublicKey        string  `json:"publicKey,omitempty"`
	PaymentReference string  `json:"paymentReference,omitempty"`
	Orders           []Order `json:"orders,omitempty"`
}
