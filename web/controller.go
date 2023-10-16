package web

import (
	"io"
	"net/http"

	"github.com/wardonne/gopi/web/context"
)

// IController interface of controller
type IController interface {
	// Init inits controller
	Init(request *context.Request)
}

// Controller basic controller implemention
type Controller struct {
	*context.Request
}

// Init inits controller
func (controller *Controller) Init(request *context.Request) {
	controller.Request = request
}

// Response returns a basic response
func (controller *Controller) Response(statusCode int, content ...any) *context.Response {
	return context.NewResponse(statusCode, content...)
}

// JSON returns a json response
func (controller *Controller) JSON(statusCode int, content ...any) *context.JSONResponse {
	return controller.Response(statusCode, content...).JSON()
}

// XML returns a xml response
func (controller *Controller) XML(statusCode int, content ...any) *context.XMLResponse {
	return controller.Response(statusCode, content...).XML()
}

// YAML returns a yaml response
func (controller *Controller) YAML(statusCode int, content ...any) *context.YAMLResponse {
	return controller.Response(statusCode, content...).YAML()
}

// TOML returns a toml response
func (controller *Controller) TOML(statusCode int, content ...any) *context.TOMLResponse {
	return controller.Response(statusCode, content...).TOML()
}

// Protobuf returns a protobuf response
func (controller *Controller) Protobuf(statusCode int, content ...any) *context.ProtobufResponse {
	return controller.Response(statusCode, content...).Protobuf()
}

// Reader returns a reader response
func (controller *Controller) Reader(statusCode int, r io.Reader) *context.ReaderResponse {
	return controller.Response(statusCode).Reader(r)
}

// Redirect returns a redirect response
func (controller *Controller) Redirect(location string, statusCode ...int) *context.RedirectResponse {
	code := http.StatusFound
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return controller.Response(code).Redirect(location)
}

// File returns a file response
func (controller *Controller) File(statusCode int, file string) *context.FileResponse {
	return controller.Response(statusCode).File(file)
}

// Stream returns a streamed response
func (controller *Controller) Stream(step func(io.Writer) bool) *context.StreamedResponse {
	return (&context.StreamedResponse{Response: &context.Response{}}).SetStep(step)
}
