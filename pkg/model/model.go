package model

import "github.com/shopspring/decimal"

type PaymentResponseData struct {
	Code          string          `json:"code"`
	Message       string          `json:"message"`
	Hash          HashData        `json:"hash,omitempty"`
	Payment       PaymentData     `json:"payments,omitempty"`
	Customer      CustomerData    `json:"customers,omitempty"`
	Subscription  Subscription    `json:"subscriptions,omitempty"`
	MerchantBanks []MerchantBanks `json:"merchantBanks,omitempty"`
	Products      []OrderProducts `json:"products,omitempty"`
}

type Subscription struct {
	PublicKey         string `json:"publicKey,omitempty"`
	Amount            int    `json:"amount,omitempty"`
	Country           string `json:"country,omitempty"`
	CustomerID        string `json:"customerId,omitempty"`
	CardName          string `json:"cardName,omitempty"`
	CardNumber        string `json:"cardNumber,omitempty"`
	Plan              string `json:"plan,omitempty"`
	Status            string `json:"status,omitempty"`
	BillingID         string `json:"billingId,omitempty"`
	AuthorizationCode string `json:"authorizationCode,omitempty"`
	StartDate         string `json:"startDate,omitempty"`
	CreatedAt         int64  `json:"createdAt,omitempty"`
}

type HashData struct {
	Hash string `json:"hash,omitempty"`
}

type OrderProducts struct {
	Amount             string `json:"amount,omitempty"`
	Currency           string `json:"currency,omitempty"`
	ProductId          string `json:"productId,omitempty"`
	ProductDescription string `json:"productDescription,omitempty"`
	PaymentReference   string `json:"paymentReference,omitempty"`
	OrderId            string `json:"orderId,omitempty"`
	CreatedAt          string `json:"createdAt,omitempty"`
	UpdatedAt          string `json:"updatedAt,omitempty"`
}

type PaymentData struct {
	RedirectLink       string           `json:"redirectLink,omitempty"`
	PaymentStatus      string           `json:"paymentStatus,omitempty"`
	PaymentReference   string           `json:"paymentReference,omitempty"`
	LinkingReference   string           `json:"linkingReference,omitempty"`
	Reference          string           `json:"reference,omitempty"`
	RedirectUrl        string           `json:"redirectUrl,omitempty"`
	Amount             string           `json:"amount,omitempty"`
	WalletName         string           `json:"walletName,omitempty"`
	BankName           string           `json:"bankName,omitempty"`
	AccountNumber      string           `json:"accountNumber,omitempty"`
	Reason             string           `json:"reason,omitempty"`
	ProcessedTime      string           `json:"transactionProcessedTime,omitempty"`
	Fee                string           `json:"fee,omitempty"`
	PublicKey          string           `json:"publicKey,omitempty"`
	Mobilenumber       string           `json:"mobileNumber,omitempty"`
	PaymentType        string           `json:"paymentType,omitempty"`
	ProductId          string           `json:"productId,omitempty"`
	ProductDescription string           `json:"productDescription,omitempty"`
	MaskedPan          string           `json:"maskedPan,omitempty"`
	GatewayMessage     string           `json:"gatewayMessage,omitempty"`
	GatewayCode        string           `json:"gatewayCode,omitempty"`
	GatewayRef         string           `json:"gatewayref,omitempty"`
	BusinessName       string           `json:"businessName,omitempty"`
	Mode               string           `json:"mode,omitempty"`
	CallbackUrl        string           `json:"callbackurl,omitempty"`
	ChannelType        string           `json:"channelType,omitempty"`
	SourceIP           string           `json:"sourceIP,omitempty"`
	DeviceType         string           `json:"deviceType,omitempty"`
	CardBin            string           `json:"cardBin,omitempty"`
	LastFourDigits     string           `json:"lastFourDigits,omitempty"`
	Country            string           `json:"country,omitempty"`
	CardToken          string           `json:"cardToken,omitempty"`
	Currency           string           `json:"currency,omitempty"`
	Status             string           `json:"status,omitempty"`
	PaymentBreakdown   PaymentBreakdown `json:"paymentBreakdown,omitempty"`
}

type PaymentBreakdown struct {
	Amount decimal.Decimal `json:"amount,omitempty"`
	Fee    decimal.Decimal `json:"fee,omitempty"`
	Total  decimal.Decimal `json:"total,omitempty"`
}

type CustomerData struct {
	CustomerID     string `json:"customerId,omitempty"`
	CustomerName   string `json:"customerName,omitempty"`
	CustomerMobile string `json:"customerMobile,omitempty"`
	CustomerEmail  string `json:"customerEmail,omitempty"`
	Fee            string `json:"fee,omitempty"`
}

type PaymentResponse struct {
	Status string              `json:"status"`
	Data   PaymentResponseData `json:"data"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type RequiredFields struct {
	AccountName   string `json:"accountName,omitempty"`
	AccountNumber string `json:"accountNumber,omitempty"`
	IsBankCode    string `json:"isBankCode,omitempty"`
	Bvn           string `json:"bvn,omitempty"`
	DateOfBirth   string `json:"dateOfBirth,omitempty"`
	MobileNumber  string `json:"mobileNumber,omitempty"`
}

type MerchantBanks struct {
	BankName       string          `json:"bankName,omitempty"`
	BankCode       string          `json:"bankCode,omitempty"`
	URL            string          `json:"url,omitempty"`
	Logo           string          `json:"logo,omitempty"`
	Status         string          `json:"status,omitempty"`
	MinimumAmount  decimal.Decimal `json:"minimumAmount,omitempty"`
	RequiredFields RequiredFields  `json:"requiredFields,omitempty"`
}

type PreauthorizationPayload struct {
	PublicKey          string `json:"publicKey,omitempty"`
	Amount             string `json:"amount,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Country            string `json:"country,omitempty"`
	PaymentReference   string `json:"paymentReference,omitempty"`
	ProductDescription string `json:"productDescription,omitempty"`
}

type CardPayload struct {
	PublicKey          string `json:"publicKey,omitempty"`
	Amount             string `json:"amount,omitempty"`
	Fee                string `json:"fee,omitempty"`
	FullName           string `json:"fullName,omitempty"`
	MobileNumber       string `json:"mobileNumber,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Country            string `json:"country,omitempty"`
	PaymentReference   string `json:"paymentReference,omitempty"`
	Email              string `json:"email,omitempty"`
	ProductId          string `json:"productId,omitempty"`
	ProductDescription string `json:"productDescription,omitempty"`
	ClientAppCode      string `json:"clientAppCode,omitempty"`
	RedirectUrl        string `json:"redirectUrl,omitempty"`
	CallbackUrl        string `json:"callbackUrl,omitempty"`
	PaymentType        string `json:"paymentType,omitempty"`
	ScheduleId         string `json:"scheduleId,omitempty"`
	ChannelType        string `json:"channelType,omitempty"`
	DeviceType         string `json:"deviceType,omitempty"`
	SourceIP           string `json:"sourceIP,omitempty"`
	CardNumber         string `json:"cardNumber,omitempty"`
	Cvv                string `json:"cvv,omitempty"`
	ExpiryMonth        string `json:"expiryMonth,omitempty"`
	ExpiryYear         string `json:"expiryYear,omitempty"`
	Pin                string `json:"pin,omitempty"`
	Source             string `json:"source,omitempty"`
	StartDate          string `json:"startDate,omitempty"`
	Retry              string `json:"retry,omitempty"`
	CardName           string `json:"cardName,omitempty"`
	Type               string `json:"type,omitempty"`
	InvoiceNumber      string `json:"invoiceNumber,omitempty"`
	IsCvv              string `json:"isCvv,omitempty"`
	PlanId             string `json:"planId,omitempty"`
	BillingCycle       string `json:"billingCycle,omitempty"`
	CustomerId         string `json:"customerId,omitempty"`
	BillingPeriod      string `json:"billingPeriod,omitempty"`
	SubscriptionAmount bool   `json:"subscriptionAmount,omitempty"`
}

type ValidationPayload struct {
	LinkingReference string `json:"linkingReference"`
	Otp              string `json:"otp"`
}

type MobileMoneyPayload struct {
	FullName           string `json:"fullName,omitempty"`
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
	ProductId          string `json:"productId"`
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

type OrderData struct {
	OrderID            string `json:"orderId,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Amount             string `json:"amount,omitempty"`
	ProductID          string `json:"productId,omitempty"`
	ProductDescription string `json:"productDescription,omitempty"`
}

type OrderPaymentPayload struct {
	Email            string      `json:"email,omitempty"`
	PublicKey        string      `json:"publicKey,omitempty"`
	PaymentReference string      `json:"paymentReference,omitempty"`
	FullName         string      `json:"fullName,omitempty"`
	OrderType        string      `json:"orderType,omitempty"`
	MobileNumber     string      `json:"mobileNumber,omitempty"`
	CallbackURL      string      `json:"callbackUrl,omitempty"`
	Country          string      `json:"country,omitempty"`
	Currency         string      `json:"currency,omitempty"`
	Amount           string      `json:"amount,omitempty"`
	Orders           []OrderData `json:"orders,omitempty"`
}

type OrderPostPaymentPayload struct {
	PublicKey        string      `json:"publicKey,omitempty"`
	PaymentReference string      `json:"paymentReference,omitempty"`
	Orders           []OrderData `json:"orders,omitempty"`
}

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
