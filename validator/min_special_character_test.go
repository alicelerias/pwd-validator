package validator

import (
	"testing"
)

func TestMinSpecialCharacter(t *testing.T) {
	minSpecialCharacter := NewMinSpecialCharacter(2)

	// test 2 Special Characters

	if valid := minSpecialCharacter.Validate("minhA1senha@["); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test Special Characters

	if valid := minSpecialCharacter.Validate("minhA1senha!@#$%^&*()-/+{}[]"); !valid {
		t.Fatal("Password did not pass, but should pass!")
	}

	// test 1 Special Characters

	if valid := minSpecialCharacter.Validate("minhA1senha!"); valid {
		t.Fatal("Password passed, but should not pass!")
	}

}
