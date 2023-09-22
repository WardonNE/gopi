package binding

import (
	"encoding/xml"
	"errors"
	"net/http"
)

type XMLParser struct {
}

func (xmlParser *XMLParser) Parse(request *http.Request, container any) error {
	if request == nil || request.Body == nil {
		return errors.New("invalid request")
	}
	decoder := xml.NewDecoder(request.Body)
	return decoder.Decode(container)
}

type xmlBinding struct {
}

func (x *xmlBinding) Parser() Parser {
	return &XMLParser{}
}
