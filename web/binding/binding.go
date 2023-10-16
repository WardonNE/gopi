package binding

import "net/http"

// Parser is the function used to parse request params into container
type Parser func(r *http.Request, container any) error

// Binding is the alias of [Parser]
type Binding = Parser
