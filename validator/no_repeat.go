package validator

const NoRepeatKey = "noRepeted"

// NoRepeat rule implements Validator
type NoRepeat struct {
}

// NewNoRepeat constructor
func NewNoRepeat() *NoRepeat {
	return &NoRepeat{}
}

// Validate returns true if the value has not a repeating character
func (s NoRepeat) Validate(value string) bool {
	var slow rune = 0
	for _, fast := range value {
		if slow == 0 {
			slow = fast
			continue
		}
		if slow == fast {
			return false
		}
		slow = fast
	}
	return true
}

func (s *NoRepeat) GetKey() string {
	return NoRepeatKey
}
