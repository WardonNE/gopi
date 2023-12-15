package validation

import "github.com/go-playground/validator/v10"

type ICustomValidateRule interface {
	Tag() string
	Validate(fl validator.FieldLevel) bool
	SkipNull() bool
}

type CustomValidation = func(form IValidateForm)
