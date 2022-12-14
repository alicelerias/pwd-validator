package validator

import (
	"testing"
)

func TestMinSize(t *testing.T) {
	minSize := NewMinSize(8)

	// test 8 chars

	if valid := minSize.Validate("00000000"); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test 9 chars

	if valid := minSize.Validate("000000001"); !valid {
		t.Fatal("Password passed, but should not pass!")
	}

	// test 7 chars
	if valid := minSize.Validate("000000"); valid {
		t.Fatal("Password did not pass, but should pass!")
	}
}
