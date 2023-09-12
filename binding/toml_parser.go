package binding

import (
	"errors"
	"net/http"

	"github.com/pelletier/go-toml/v2"
)

type TOMLParser struct{}

func (tomlParser *TOMLParser) Parse(request *http.Request, container any) error {
	if request == nil || request.Body == nil {
		return errors.New("invalid request")
	}
	decoder := toml.NewDecoder(request.Body)
	return decoder.Decode(container)
}
