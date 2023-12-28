package validation

import "github.com/go-playground/validator/v10"

// ICustomValidateRule custom validation tag interface
type ICustomValidateRule interface {
	Tag() string
	Validate(fl validator.FieldLevel) bool
	SkipNull() bool
}

// CustomValidation custom validation function
type CustomValidation = func(form IValidateForm)
