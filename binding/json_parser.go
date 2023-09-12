package binding

import (
	"encoding/json"
	"errors"
	"net/http"
)

type JSONParser struct {
}

func (jsonParser *JSONParser) Parse(request *http.Request, container any) error {
	if request == nil || request.Body == nil {
		return errors.New("invalid request")
	}
	decoder := json.NewDecoder(request.Body)
	return decoder.Decode(container)
}
