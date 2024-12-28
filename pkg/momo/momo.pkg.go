package momo

type MobileMoneyPayload struct {
	Fullname           string `json:"fullName,omitempty"`
	Email              string `json:"email,omitempty"`
	MobileNumber       string `json:"mobileNumber,omitempty"`
	PublicKey          string `json:"publicKey,omitempty"`
	PaymentReference   string `json:"paymentReference,omitempty"`
	DeviceType         string `json:"deviceType,omitempty"`
	SourceIP           string `json:"sourceIP,omitempty"`
	Currency           string `json:"currency,omitempty"`
	ProductDescription string `json:"productDescription,omitempty"`
	Country            string `json:"country,omitempty"`
	Network            string `json:"network,omitempty"`
	VoucherCode        string `json:"voucherCode,omitempty"`
	Fee                string `json:"fee,omitempty"`
	Amount             string `json:"amount,omitempty"`
	ProductId          string `json:"productId,omitempty"`
	PaymentType        string `json:"paymentType,omitempty"`
}
