package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate = validator.New()

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)

	if err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}

func ValidateVar(v any, reqs string) error {
	err := validate.Var(v, reqs)

	if err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}
