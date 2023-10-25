package contract

import (
	"net/http"

	"github.com/wardonne/gopi/web/context"
)

type ErrorHandler interface {
	Render(*http.Request, error) context.IResponse
}
