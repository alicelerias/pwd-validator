package validator

import (
	"errors"
	"fmt"

	"github.com/alicelerias/pwd-validator/graph/model"
)

// Validator interface
type Validator interface {
	Validate(value string) bool
	GetKey() string
}

// getValidator return a Validator for the model.Rule
// If the model.Rule is not found returns error
func getValidator(rule *model.Rule) (Validator, error) {
	switch rule.Name {
	case "minSize":
		return NewMinSize(rule.Value), nil

	case "minDigit":
		return NewMinDigit(rule.Value), nil

	case "minLowercase":
		return NewMinLowercase(rule.Value), nil

	case "minUppercase":
		return NewMinUppercase(rule.Value), nil

	case "minSpecialChar":
		return NewMinSpecialCharacter(rule.Value), nil

	case "noRepeted":
		return NewNoRepeat(), nil
	}
	return nil, errors.New("Invalid rule!")
}

// Validate valida password para varias rules
// returns true if all rules passes or else
// list of broken rules
func Validate(password *string, rules []*model.Rule) (bool, []string) {
	noMatch := []string{}

	for _, rule := range rules {
		validator, err := getValidator(rule)
		if err != nil {
			fmt.Print("Non-existent validator!!")
			continue
		}

		if !validator.Validate(*password) {
			noMatch = append(noMatch, validator.GetKey())

		}

	}
	return len(noMatch) == 0, noMatch
}
