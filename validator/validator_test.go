package validator

import (
	"testing"

	"github.com/alicelerias/pwd-validator/graph/model"
)

func TestValidator(t *testing.T) {
	password := "MiNhaSenhaForte@!891"
	wrongPassword := "test"
	rules := []*model.Rule{
		{
			Name:  "minDigit",
			Value: 3,
		},

		{
			Name:  "minUppercase",
			Value: 2,
		},

		{
			Name:  "minLowercase",
			Value: 2,
		},

		{
			Name:  "minSpecialChar",
			Value: 2,
		},

		{
			Name:  "minSize",
			Value: 8,
		},

		{
			Name:  "noRepeted",
			Value: 0,
		},
	}

	if valid, noMatch := Validate(&password, rules); !valid || len(noMatch) >= 1 {
		t.Fatal("Password not passed, but should pass!")
	}

	if valid, noMatch := Validate(&wrongPassword, rules); valid || len(noMatch) == 0 {
		t.Fatal("Wrong password passed, should not pass!")
	}
}
