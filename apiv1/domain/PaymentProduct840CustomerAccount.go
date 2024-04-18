// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProduct840CustomerAccount represents class PaymentProduct840CustomerAccount
type PaymentProduct840CustomerAccount struct {
	AccountID             *string `json:"accountId,omitempty"`
	BillingAgreementID    *string `json:"billingAgreementId,omitempty"`
	CompanyName           *string `json:"companyName,omitempty"`
	ContactPhone          *string `json:"contactPhone,omitempty"`
	CountryCode           *string `json:"countryCode,omitempty"`
	CustomerAccountStatus *string `json:"customerAccountStatus,omitempty"`
	CustomerAddressStatus *string `json:"customerAddressStatus,omitempty"`
	FirstName             *string `json:"firstName,omitempty"`
	PayerID               *string `json:"payerId,omitempty"`
	Surname               *string `json:"surname,omitempty"`
}

// NewPaymentProduct840CustomerAccount constructs a new PaymentProduct840CustomerAccount instance
func NewPaymentProduct840CustomerAccount() *PaymentProduct840CustomerAccount {
	return &PaymentProduct840CustomerAccount{}
}
