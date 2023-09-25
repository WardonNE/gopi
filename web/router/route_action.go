package router

import (
	"fmt"
	"reflect"

	"github.com/wardonne/gopi/pipeline"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web"
	"github.com/wardonne/gopi/web/binding"
	"github.com/wardonne/gopi/web/context"
	"github.com/wardonne/gopi/web/middleware"
)

type RouteAction struct {
	Route
	handler        string
	controller     web.IController
	controllerType reflect.Type
}

func (action *RouteAction) AS(name string) IRoute {
	action.name = name
	return action
}

func (action *RouteAction) Use(middlewares ...middleware.IMiddleware) IRoute {
	action.middlewares.AddAll(middlewares...)
	return action
}

func (action *RouteAction) Validate(form validation.IValidateForm, bindings ...binding.Binding) IRoute {
	formType := reflect.TypeOf(form)
	if formType.Kind() != reflect.Ptr {
		panic("Non-ptr: " + formType.String())
	}
	action.validation = middleware.Validation(form, bindings...)
	return action
}

func (action *RouteAction) Handler() string {
	return fmt.Sprintf("(%s).%s", action.controllerType.String(), action.handler)
}

func (action *RouteAction) HandleRequest(request *context.Request) context.IResponse {
	var controllerValue reflect.Value
	if action.controllerType.Kind() == reflect.Ptr {
		controllerValue = reflect.New(action.controllerType.Elem())
	} else {
		controllerValue = reflect.New(action.controllerType)
	}
	controllerValue.MethodByName("Init").Call([]reflect.Value{
		reflect.ValueOf(request),
	})
	pl := new(pipeline.Pipeline[*context.Request, context.IResponse])
	pipes := make([]pipeline.IPipe[*context.Request, context.IResponse], 0, action.middlewares.Count())
	action.middlewares.Range(func(middleware middleware.IMiddleware) bool {
		pipes = append(pipes, pipeline.AsPipe[*context.Request, context.IResponse](middleware))
		return true
	})
	pl = pl.Send(request).Through(pipes...)
	if action.HasValidation() {
		pl = pl.AppendThrough(action.validation)
	}
	return pl.Then(func(request *context.Request) context.IResponse {
		outputs := controllerValue.MethodByName(action.handler).Call([]reflect.Value{})
		resp := outputs[0].Interface().(context.IResponse)
		return resp
	})
}
