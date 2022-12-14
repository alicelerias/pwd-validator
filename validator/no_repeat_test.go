package validator

import (
	"testing"
)

func TestNoRepeat(t *testing.T) {
	noRepeat := NewNoRepeat()

	// test no repeating
	if valid := noRepeat.Validate("mInha1senha"); !valid {
		t.Fatal("Password not passed, but should pass!")
	}

	// test repeating
	if valid := noRepeat.Validate("minH1asenhaa"); valid {
		t.Fatal("Password passed, but should not pass!")
	}

}
