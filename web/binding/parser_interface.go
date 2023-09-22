package binding

import "net/http"

type Parser interface {
	Parse(request *http.Request, container any) error
}
