package middleware

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/wardonne/gopi/binding"
	"github.com/wardonne/gopi/context"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web"
)

type ValidateMiddleware struct {
	form      validation.IValidateForm
	formType  reflect.Type
	validator *validation.Validator
	bindings  []binding.Binding
}

func (v *ValidateMiddleware) Handle(request *context.Request, next web.Handler) context.IResponse {
	form := reflect.New(v.formType).Interface().(validation.IValidateForm)
	if err := request.Bind(form, v.bindings...); err != nil {
		panic(err)
	}
	translator := v.validator.Translator(request.GetString("language", "en"))
	if err := v.validator.Struct(form); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			if translator != nil {
				messages := errs.Translate(translator)
				for key, message := range messages {
					form.AddError(key, message)
				}
			} else {
				for _, err := range errs {
					form.AddError(err.Field(), err.Error())
				}
			}
		} else {
			panic(err)
		}
	}
	return next(request)
}

func Validation(form validation.IValidateForm, bindings ...binding.Binding) *ValidateMiddleware {
	vm := &ValidateMiddleware{
		form:      form,
		formType:  reflect.TypeOf(form).Elem(),
		validator: validation.Default(),
		bindings:  bindings,
	}
	return vm
}
