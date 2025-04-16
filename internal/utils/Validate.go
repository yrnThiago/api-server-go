package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateStruct(input any) error {
	var customError string
	validate = validator.New()

	if err := validate.Struct(input); err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				customError += fmt.Sprintf("%s: %s; ", e.Field(), e.Tag())
			}
		}

		return validateErrs
	}

	return nil
}
