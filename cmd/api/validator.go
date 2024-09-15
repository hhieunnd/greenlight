package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func (app *application) validate(input any) (err map[string]string) {
	validationErr := app.validator.Struct(input)

	errors := make(map[string]string)

	if validationErr != nil {
		for _, fieldError := range validationErr.(validator.ValidationErrors) {
			fieldName := fieldError.Field()
			tag := fieldError.Tag()
			param := fieldError.Param()

			var message string
			switch tag {
			case "required":
				message = fmt.Sprintf("Field '%s' is required.", fieldName)
			case "email":
				message = fmt.Sprintf("Field '%s' must be a valid email address.", fieldName)
			case "gte":
				message = fmt.Sprintf("Field '%s' must be greater than or equal to %s.", fieldName, param)
			default:
				message = fmt.Sprintf("Field '%s' failed validation for '%s' tag.", fieldName, tag)
			}

			errors[fieldName] = message
		}

		return errors
	}

	return nil
}
