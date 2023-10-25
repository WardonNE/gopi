package context

import (
	"net/http"

	"github.com/pelletier/go-toml/v2"
)

// TOMLResponse used to sends TOML-encoded data
type TOMLResponse struct {
	*Response
	data any
}

// SetContent sets response body content
func (tomlResponse *TOMLResponse) SetContent(data any) {
	tomlResponse.data = data
}

// Send sends the response
func (tomlResponse *TOMLResponse) Send(w http.ResponseWriter, r *http.Request) {
	tomlBytes, err := toml.Marshal(tomlResponse.data)
	if err != nil {
		panic(err)
	}
	tomlResponse.content = tomlBytes
	tomlResponse.SetHeader("Content-Type", MIMETOML)
	tomlResponse.Response.Send(w, r)
}
