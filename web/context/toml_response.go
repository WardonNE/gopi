package context

import (
	"net/http"

	"github.com/pelletier/go-toml/v2"
)

type TOMLResponse struct {
	*Response
	data any
}

func (tomlResponse *TOMLResponse) SetContent(data any) IResponse {
	tomlResponse.data = data
	return tomlResponse
}

func (tomlResponse *TOMLResponse) Send(w http.ResponseWriter, r *http.Request) {
	tomlBytes, err := toml.Marshal(tomlResponse.data)
	if err != nil {
		panic(err)
	}
	tomlResponse.content = tomlBytes
	tomlResponse.SetHeader("Content-Type", MIMETOML)
	tomlResponse.Response.Send(w, r)
}
