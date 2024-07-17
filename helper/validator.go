package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct[T any](s T) error {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var errMsg string
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			errMsg += fmt.Sprintf("Failed to validate %s: %s\n", field, err.Error())
		}
		return fmt.Errorf("error validating struct: %s", errMsg)
	}
	return nil
}
