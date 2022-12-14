package validator

import (
	"testing"
)

func TestMinUppercase(t *testing.T) {
	minUppercase := NewMinUppercase(2)

	// test 2 Uppercase

	if valid := minUppercase.Validate("miNhaPalavra"); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test 4 Uppercase

	if valid := minUppercase.Validate("miNhaPalaVRa"); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test 1 Uppercase
	if valid := minUppercase.Validate("minhaPalavra"); valid {
		t.Fatal("Password passed, but should not pass!")
	}

	// test number 2 Uppercase

	if valid := minUppercase.Validate("miNha2Palavra1"); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test number 1 Uppercase

	if valid := minUppercase.Validate("minha2Palavra1"); valid {
		t.Fatal("Password passed, but should not pass!")
	}
}
