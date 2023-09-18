package binding

import (
	"mime/multipart"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wardonne/gopi/utils"
)

type URIParser struct {
	params httprouter.Params
}

func (uriParser *URIParser) Parse(request *http.Request, container any) error {
	form := &multipart.Form{
		Value: make(map[string][]string),
		File:  make(map[string][]*multipart.FileHeader),
	}
	for _, param := range uriParser.params {
		form.Value[param.Key] = []string{param.Value}
	}
	return utils.FormDataToStruct(form, container, "param")
}

type uriBinding struct {
}

func (u *uriBinding) Parser() Parser {
	return &URIParser{}
}
