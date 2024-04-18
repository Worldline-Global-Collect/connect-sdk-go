// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// CustomerAccount represents class CustomerAccount
type CustomerAccount struct {
	Authentication                *CustomerAccountAuthentication `json:"authentication,omitempty"`
	ChangeDate                    *string                        `json:"changeDate,omitempty"`
	ChangedDuringCheckout         *bool                          `json:"changedDuringCheckout,omitempty"`
	CreateDate                    *string                        `json:"createDate,omitempty"`
	HadSuspiciousActivity         *bool                          `json:"hadSuspiciousActivity,omitempty"`
	HasForgottenPassword          *bool                          `json:"hasForgottenPassword,omitempty"`
	HasPassword                   *bool                          `json:"hasPassword,omitempty"`
	PasswordChangeDate            *string                        `json:"passwordChangeDate,omitempty"`
	PasswordChangedDuringCheckout *bool                          `json:"passwordChangedDuringCheckout,omitempty"`
	PaymentAccountOnFile          *PaymentAccountOnFile          `json:"paymentAccountOnFile,omitempty"`
	PaymentAccountOnFileType      *string                        `json:"paymentAccountOnFileType,omitempty"`
	PaymentActivity               *CustomerPaymentActivity       `json:"paymentActivity,omitempty"`
}

// NewCustomerAccount constructs a new CustomerAccount instance
func NewCustomerAccount() *CustomerAccount {
	return &CustomerAccount{}
}
