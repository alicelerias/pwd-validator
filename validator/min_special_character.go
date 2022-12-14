package validator

import "unicode"

const MinSpecialCharacterKey = "minSpecialChar"

// MinSpecialCharacter rule implements Validator
type MinSpecialCharacter struct {
	n int
}

// MewMinSpecialCharacter constructor
func NewMinSpecialCharacter(n int) *MinSpecialCharacter {
	return &MinSpecialCharacter{n}
}

// Validate returns true if the value has a minimum special characters
func (s MinSpecialCharacter) Validate(value string) bool {
	return minValidChar(s.n, value, func(r rune) bool {
		return !unicode.IsDigit(r) && !unicode.IsLetter(r)
	})
}

func (s *MinSpecialCharacter) GetKey() string {
	return MinSpecialCharacterKey
}
