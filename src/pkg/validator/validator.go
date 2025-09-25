package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(data interface{}, messages map[string]string) []*ValidationError {
	var errors []*ValidationError

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			key := fmt.Sprintf("%s.%s", err.Field(), err.Tag())
			msg, ok := messages[key]
			if !ok {
				msg = fmt.Sprintf("%s is %s", err.Field(), err.Tag())
			}
			errors = append(errors, &ValidationError{
				Field:   err.Field(),
				Message: msg,
			})
		}
	}

	return errors
}
