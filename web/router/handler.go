package router

import (
	"net/http"

	"github.com/wardonne/gopi/web/context"
)

// Handler http handler
type Handler = func(*context.Request) context.IResponse

func defaultErrorHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
}
