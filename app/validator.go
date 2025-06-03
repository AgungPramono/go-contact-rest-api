package app

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func ParseValidationErrors(err error) string {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldErr := range validationErrors {
			fieldName := fieldErr.Field()
			switch fieldErr.Tag() {
			case "required":
				return fieldName + " wajib diisi"
			case "min":
				return fieldName + " minimal " + fieldErr.Param() + " karakter"
			case "max":
				return fieldName + " maksimal " + fieldErr.Param() + " karakter"
			default:
				return "Field " + fieldName + " tidak valid"
			}
		}
	}
	return err.Error()
}
