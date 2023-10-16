package binding

import (
	"encoding/xml"
	"errors"
	"net/http"
)

// XML implements [Parser], it parses request body into container
// Make sure your request body is not nil and is XML-Encoded, or it will returns an error
//
// Example:
//
//	var container = &struct{
//	    Name string `xml:"name"`
//	    Age int `xml:"age"`
//	    Tags []string `xml:"tags"`
//	    VIP bool `xml:"vip"`
//	}{}
//
//	err := XML(request, container)
//	if err != nil {
//	    panic(err)
//	}
func XML(request *http.Request, container any) error {
	if request == nil || request.Body == nil {
		return errors.New("invalid request")
	}
	decoder := xml.NewDecoder(request.Body)
	return decoder.Decode(container)
}
