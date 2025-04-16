package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateStruct(input any) error {
	validate = validator.New()

	if err := validate.Struct(input); err != nil {
		return err
	}

	return nil
}
