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
	Payload       []Payload       `json:"payload,omitempty"`
}

type Payload struct {
	ID                        int             `json:"id,omitempty"`
	FullName                  string          `json:"fullName,omitempty"`
	PublicKey                 string          `json:"publicKey,omitempty"`
	Processor                 string          `json:"processor,omitempty"`
	PaymentType               string          `json:"paymentType,omitempty"`
	ShopperReference          string          `json:"shopperReference,omitempty"`
	Amount                    decimal.Decimal `json:"amount,omitempty"`
	ProductID                 string          `json:"productId,omitempty"`
	ProductDescription        string          `json:"productDescription,omitempty"`
	Email                     string          `json:"email,omitempty"`
	Quantity                  int             `json:"quantity,omitempty"`
	MaskedPan                 string          `json:"maskedPan,omitempty"`
	Gateway                   string          `json:"gateway,omitempty"`
	GatewayMessage            string          `json:"gatewayMessage,omitempty"`
	GatewayCode               string          `json:"gatewayCode,omitempty"`
	TransactionRef            string          `json:"transactionRef,omitempty"`
	GatewayRef                string          `json:"gatewayRef,omitempty"`
	BusinessName              string          `json:"businessName,omitempty"`
	Fee                       decimal.Decimal `json:"fee,omitempty"`
	Mode                      string          `json:"mode,omitempty"`
	CallbackUrl               string          `json:"callbackUrl,omitempty"`
	RedirectUrl               string          `json:"redirectUrl,omitempty"`
	ChannelType               string          `json:"channelType,omitempty"`
	PaymentKey                string          `json:"paymentKey,omitempty"`
	SourceIP                  string          `json:"sourceIP,omitempty"`
	DeviceType                string          `json:"deviceType,omitempty"`
	ClientAppCode             string          `json:"clientAppCode,omitempty"`
	CardBin                   string          `json:"cardBin,omitempty"`
	LastFourDigits            string          `json:"lastFourDigits,omitempty"`
	Type                      string          `json:"type,omitempty"`
	LinkingReference          string          `json:"linkingreference,omitempty"`
	Country                   string          `json:"country,omitempty"`
	Currency                  string          `json:"currency,omitempty"`
	SmsProvider               string          `json:"smsProvider,omitempty"`
	CustomerID                string          `json:"customerId,omitempty"`
	InternalReference         string          `json:"internalreference,omitempty"`
	AccountNumber             string          `json:"accountNumber,omitempty"`
	Narration                 string          `json:"narration,omitempty"`
	CreditAccountName         string          `json:"creditAccountName,omitempty"`
	TransferType              string          `json:"transferType,omitempty"`
	PaymentReference          string          `json:"paymentReference,omitempty"`
	BatchID                   string          `json:"batchId,omitempty"`
	SessionID                 string          `json:"sessionId,omitempty"`
	BankName                  string          `json:"bankName,omitempty"`
	CreditAccountNumber       string          `json:"creditAccountNumber,omitempty"`
	BankCode                  string          `json:"bankCode,omitempty"`
	AlternatePaymentReference string          `json:"alternatePaymentReference,omitempty"`
	SettlementCode            string          `json:"settlementCode,omitempty"`
	SettlementMessage         string          `json:"settlementMessage,omitempty"`
	SettlementTime            string          `json:"settlementTime,omitempty"`
	OrderStatusCode           string          `json:"orderStatusCode,omitempty"`
	OrderStatusMessage        string          `json:"orderStatusMessage,omitempty"`
	Status                    string          `json:"status,omitempty"`
	MobileNumber              string          `json:"mobileNumber,omitempty"`
	DateOfBirth               string          `json:"dateOfBirth,omitempty"`
	BranchPhoneNumber         string          `json:"branchPhoneNumber,omitempty"`
	TransferredAmount         float64         `json:"transferedAmount,omitempty"`
	ScheduleID                string          `json:"scheduleId,omitempty"`
	IsCardInternational       string          `json:"isCardInternational,omitempty"`
	Reason                    string          `json:"reason,omitempty"`
	Retry                     bool            `json:"retry,omitempty"`
	MetaData                  any             `json:"metaData,omitempty"`
	Event                     []any           `json:"event,omitempty"`
	Order                     []any           `json:"order,omitempty"`
	CreatedAt                 string          `json:"createdAt,omitempty"`
	UpdatedAt                 string          `json:"updatedAt,omitempty"`
	CardName                  string          `json:"cardName,omitempty"`
	IsNigerianCard            string          `json:"isNigeriancard,omitempty"`
	CardCountry               string          `json:"cardCountry,omitempty"`
	IntCurrency               string          `json:"intCurrency,omitempty"`
	Rate                      decimal.Decimal `json:"rate,omitempty"`
	InCardProcessingFee       decimal.Decimal `json:"inCardProcessingFee,omitempty"`
	IntAmountCharge           decimal.Decimal `json:"intAmountCharge,omitempty"`
	ProcessorCode             string          `json:"processorCode,omitempty"`
	ProcessorMessage          string          `json:"processorMessage,omitempty"`
	InvoiceNumber             string          `json:"invoiceNumber,omitempty"`
	BillID                    string          `json:"billId,omitempty"`
	LocationPhoneNumber       string          `json:"locationPhoneNumber,omitempty"`
	PocketReferenceID         string          `json:"pocketReferenceId,omitempty"`
	TransferAccountType       string          `json:"transferAccountType,omitempty"`
	Bearer                    string          `json:"bearer,omitempty"`
	TransLink                 string          `json:"transLink,omitempty"`
	VendorID                  string          `json:"vendorId,omitempty"`
	PayLinkEnvironment        string          `json:"payLinkEnvironment,omitempty"`
	PayLinkStatus             string          `json:"payLinkStatus,omitempty"`
	PayLinkAmount             decimal.Decimal `json:"payLinkAmount,omitempty"`
	PayLinkAdditionalData     string          `json:"payLinkAdditionalData,omitempty"`
	PayLinkName               string          `json:"payLinkName,omitempty"`
	PayLinkDescription        string          `json:"payLinkDescription,omitempty"`
	PayLinkRedirectUrl        string          `json:"payLinkRedirectUrl,omitempty"`
	PayLinkSuccessMessage     string          `json:"payLinkSuccessMessage,omitempty"`
	PaymentLinkID             string          `json:"paymentLinkId,omitempty"`
	PayLinkCustomizationName  string          `json:"payLinkCustomizationName,omitempty"`
	PayLinkFrequency          string          `json:"payLinkFrequency,omitempty"`
	PayLinkIsOneTimeUse       bool            `json:"payLinkIsOneTimeUse,omitempty"`
	TerminalID                string          `json:"terminalId,omitempty"`
	Stan                      string          `json:"stan,omitempty"`
	TransactionComplete       bool            `json:"transactionComplete,omitempty"`
	CardExpiryMonth           string          `json:"cardExpiryMonth,omitempty"`
	CardExpiryYear            string          `json:"cardExpiryYear,omitempty"`
	Tokenize                  bool            `json:"tokenize,omitempty"`
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
	AuthorizationCode  string           `json:"authorizationCode,omitempty"`
	WalletName         string           `json:"walletName,omitempty"`
	Wallet             string           `json:"wallet,omitempty"`
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
	Status  string              `json:"status,omitempty"`
	Message string              `json:"message,omitempty"`
	Error   string              `json:"error,omitempty"`
	Data    PaymentResponseData `json:"data,omitempty"`
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

type ChargeToken struct {
	PublicKey         string `json:"publicKey,omitempty"`
	Amount            string `json:"amount,omitempty"`
	PaymentReference  string `json:"paymentReference,omitempty"`
	AuthorizationCode string `json:"authorizationCode,omitempty"`
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
	Tokenize           string `json:"tokenize,omitempty"`
}

type VirtualAccountPayload struct {
	PublicKey              string `json:"publicKey,omitempty"`
	FullName               string `json:"fullName,omitempty"`
	BankVerificationNumber string `json:"bankVerificationNumber,omitempty"`
	Currency               string `json:"currency,omitempty"`
	Country                string `json:"country,omitempty"`
	Reference              string `json:"reference,omitempty"`
	Email                  string `json:"email,omitempty"`
}
