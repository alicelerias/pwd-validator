package validator

const MinSizeKey = "minSize"

// MinSize rule implements Validator
type MinSize struct {
	n int
}

// NewMinSize constructor
func NewMinSize(n int) *MinSize {
	return &MinSize{n}
}

// Validate returns true if the value has a miminum len
func (s MinSize) Validate(value string) bool {
	return len(value) >= s.n
}

func (s *MinSize) GetKey() string {
	return MinSizeKey
}
