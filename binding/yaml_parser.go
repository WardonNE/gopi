package binding

import (
	"errors"
	"net/http"

	"gopkg.in/yaml.v3"
)

type YAMLParser struct{}

func (yamlParser *YAMLParser) Parse(request *http.Request, container any) error {
	if request == nil || request.Body == nil {
		return errors.New("invalid request")
	}
	decoder := yaml.NewDecoder(request.Body)
	return decoder.Decode(container)
}
