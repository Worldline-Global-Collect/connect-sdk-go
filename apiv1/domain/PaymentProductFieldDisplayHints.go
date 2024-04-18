// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// PaymentProductFieldDisplayHints represents class PaymentProductFieldDisplayHints
type PaymentProductFieldDisplayHints struct {
	AlwaysShow         *bool                           `json:"alwaysShow,omitempty"`
	DisplayOrder       *int32                          `json:"displayOrder,omitempty"`
	FormElement        *PaymentProductFieldFormElement `json:"formElement,omitempty"`
	Label              *string                         `json:"label,omitempty"`
	Link               *string                         `json:"link,omitempty"`
	Mask               *string                         `json:"mask,omitempty"`
	Obfuscate          *bool                           `json:"obfuscate,omitempty"`
	PlaceholderLabel   *string                         `json:"placeholderLabel,omitempty"`
	PreferredInputType *string                         `json:"preferredInputType,omitempty"`
	Tooltip            *PaymentProductFieldTooltip     `json:"tooltip,omitempty"`
}

// NewPaymentProductFieldDisplayHints constructs a new PaymentProductFieldDisplayHints instance
func NewPaymentProductFieldDisplayHints() *PaymentProductFieldDisplayHints {
	return &PaymentProductFieldDisplayHints{}
}
