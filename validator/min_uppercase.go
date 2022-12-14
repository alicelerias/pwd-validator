package validator

import "unicode"

const MinUppercaseKey = "minUppercase"

// MinUppercase rule implements Validator
type MinUppercase struct {
	n int
}

// NewMinUppercase  constructor
func NewMinUppercase(n int) *MinUppercase {
	return &MinUppercase{n}
}

// Validate returns true if the value has a minimum uppercase character
func (s MinUppercase) Validate(value string) bool {
	return minValidChar(s.n, value, func(r rune) bool {
		return unicode.IsUpper(r)
	})
}

func (s *MinUppercase) GetKey() string {
	return MinUppercaseKey
}
