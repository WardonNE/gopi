package validate

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web/binding"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

// ValidationEngine validation engine interface
type ValidationEngine interface {
	Translator(locale string) ut.Translator
	Struct(any) error
}

// New creates a new validation middleware instance
func New(engine ValidationEngine, form validation.IValidateForm, bindings ...binding.Binding) middleware.IMiddleware {
	formType := reflect.TypeOf(form).Elem()
	return func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		form := reflect.New(formType).Interface().(validation.IValidateForm)
		if err := request.Bind(form, bindings...); err != nil {
			panic(err)
		}
		translator := engine.Translator(*request.GetString("language", "en"))
		if err := engine.Struct(form); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				if translator != nil {
					for _, err := range errs {
						message := err.Translate(translator)
						form.AddError(err.Field(), message)
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
		customValidations := form.CustomValidations()
		for _, customValidation := range customValidations {
			customValidation(form)
		}
		return next(request)
	}
}
