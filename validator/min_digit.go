package validator

import "unicode"

var mindigitValidatorCheck Validator = &MinDigit{}

const MinDigitKey = "minDigit"

// MinDigit rule implements Validator
type MinDigit struct {
	n int
}

// NewMinDigit constructor
func NewMinDigit(n int) *MinDigit {
	return &MinDigit{n}
}

// Validate returns true if the value has a minimum digit character
func (s MinDigit) Validate(value string) bool {
	return minValidChar(s.n, value, func(r rune) bool {
		return unicode.IsDigit(r)
	})
}

func (s *MinDigit) GetKey() string {
	return MinDigitKey
}
