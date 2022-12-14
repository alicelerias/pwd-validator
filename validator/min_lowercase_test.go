package validator

import (
	"testing"
)

func TestMinLowercase(t *testing.T) {
	minLowercase := NewMinLowercase(2)

	// test 2 Lowercase

	if valid := minLowercase.Validate("MInHApALAVRA"); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test 4 Lowercase

	if valid := minLowercase.Validate("MinHaPALaVRA"); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test 1 Lowercase
	if valid := minLowercase.Validate("mINHAPALAVRA"); valid {
		t.Fatal("Password passed, but should not pass!")
	}

	// test number 2 Lowercase

	if valid := minLowercase.Validate("MInHApA2LAVRA1"); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test number 1 Uppercase

	if valid := minLowercase.Validate("MINHA2PaLAVRA1"); valid {
		t.Fatal("Password passed, but should not pass!")
	}
}
