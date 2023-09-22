package web

import (
	"io"
	"net/http"

	"github.com/wardonne/gopi/web/context"
)

type IController interface {
	Init(request *context.Request)
}

type Controller struct {
	*context.Request
}

func (controller *Controller) Init(request *context.Request) {
	controller.Request = request
}

func (controller *Controller) Response(statusCode int, content ...any) *context.Response {
	return context.NewResponse(statusCode, content...)
}

func (controller *Controller) JSON(statusCode int, content ...any) *context.JSONResponse {
	return controller.Response(statusCode, content...).JSON()
}

func (controller *Controller) XML(statusCode int, content ...any) *context.XMLResponse {
	return controller.Response(statusCode, content...).XML()
}

func (controller *Controller) YAML(statusCode int, content ...any) *context.YAMLResponse {
	return controller.Response(statusCode, content...).YAML()
}

func (controller *Controller) TOML(statusCode int, content ...any) *context.TOMLResponse {
	return controller.Response(statusCode, content...).TOML()
}

func (controller *Controller) Protobuf(statusCode int, content ...any) *context.ProtobufResponse {
	return controller.Response(statusCode, content...).Protobuf()
}

func (controller *Controller) Reader(statusCode int, r io.Reader) *context.ReaderResponse {
	return controller.Response(statusCode).Reader(r)
}

func (controller *Controller) Redirect(location string, statusCode ...int) *context.RedirectResponse {
	code := http.StatusFound
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return controller.Response(code).Redirect(location)
}

func (controller *Controller) File(statusCode int, file string) *context.FileResponse {
	return controller.Response(statusCode).File(file)
}

func (controller *Controller) Stream(step func(io.Writer) bool) *context.StreamedResponse {
	return (&context.StreamedResponse{Response: &context.Response{}}).SetStep(step)
}
