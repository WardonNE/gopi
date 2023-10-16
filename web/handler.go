package web

import (
	"github.com/wardonne/gopi/web/context"
)

// Handler http handler
type Handler = func(*context.Request) context.IResponse
