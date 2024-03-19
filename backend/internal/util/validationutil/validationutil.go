package validationutil

import (
	"sharedlambdacode/internal/excep"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(structToValidate interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(structToValidate); err != nil {
		return excep.NewInvalidExcep(err.Error())
	}

	return nil
}
