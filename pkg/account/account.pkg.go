package account

type AccountPayload struct {
	PublicKey          string `json:"publicKey"`
	Amount             string `json:"amount"`
	Fee                string `json:"fee"`
	FullName           string `json:"fullName"`
	MobileNumber       string `json:"mobileNumber"`
	Currency           string `json:"currency"`
	Country            string `json:"country"`
	PaymentReference   string `json:"paymentReference"`
	Email              string `json:"email"`
	ProductID          string `json:"productId"`
	ProductDescription string `json:"productDescription"`
	ClientAppCode      string `json:"clientAppCode"`
	ChannelType        string `json:"channelType"`
	RedirectUrl        string `json:"redirectUrl"`
	CallbackUrl        string `json:"callbackUrl"`
	PaymentType        string `json:"paymentType"`
	DeviceType         string `json:"deviceType"`
	SourceIP           string `json:"sourceIP"`
	AccountName        string `json:"accountName"`
	AccountNumber      string `json:"accountNumber"`
	BankCode           string `json:"bankCode"`
	BVN                string `json:"bvn"`
	DateOfBirth        string `json:"dateOfBirth"`
	Retry              string `json:"retry"`
	InvoiceNumber      string `json:"invoiceNumber"`
}
