package validator

import "unicode"

const MinLowercaseKey = "minLowerCase"

// MinLowerCase rule implements Validator
type MinLowercase struct {
	n int
}

// NewMinLowercase constructor
func NewMinLowercase(n int) *MinLowercase {
	return &MinLowercase{n}
}

// Validate returns true if the value has a mimimum lowercase character
func (s MinLowercase) Validate(value string) bool {
	return minValidChar(s.n, value, func(r rune) bool {
		return unicode.IsLower(r)
	})
}

func (s *MinLowercase) GetKey() string {
	return MinLowercaseKey
}
