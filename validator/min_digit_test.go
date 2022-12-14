package validator

import (
	"testing"
)

func TestMinDigit(t *testing.T) {
	minDigit := NewMinDigit(2)

	// test 2 Digits

	if valid := minDigit.Validate("Minhasenha12"); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test 4 Digits

	if valid := minDigit.Validate("miNhasenha1234"); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test 1 Digit

	if valid := minDigit.Validate("minhasenha1"); valid {
		t.Fatal("Password passed, but should not pass!")
	}

}
