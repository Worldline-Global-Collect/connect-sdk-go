// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// LengthValidator represents class LengthValidator
type LengthValidator struct {
	MaxLength *int32 `json:"maxLength,omitempty"`
	MinLength *int32 `json:"minLength,omitempty"`
}

// NewLengthValidator constructs a new LengthValidator instance
func NewLengthValidator() *LengthValidator {
	return &LengthValidator{}
}
