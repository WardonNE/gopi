package binding

import "net/http"

type IParser interface {
	Parse(request *http.Request, container any) error
}
