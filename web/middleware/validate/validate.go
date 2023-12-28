package validate

import (
	"reflect"

	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web/binding"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

// New creates a new validation middleware instance
func New(f validation.IValidateForm, bindings ...binding.Binding) middleware.IMiddleware {
	formType := reflect.TypeOf(f).Elem()
	return func(request *context.Request, next pipeline.Next[*context.Request, context.IResponse]) context.IResponse {
		form := reflect.New(formType).Interface().(validation.IValidateForm)
		if err := request.Bind(form, bindings...); err != nil {
			panic(err)
		}
		form.SetEngine(f.Engine())
		locale := request.Header("Accept-Language", "en")
		form.SetLocale(locale.String())
		if form.AutoValidate() {
			if form.BeforeValidate() {
				form.Validate(form)
			} else {
				form.AddError("onBeforeValidate", "BeforeValidate returned false")
			}
		}
		return next(request)
	}
}
