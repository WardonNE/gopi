package binding

import (
	"encoding/json"
	"errors"
	"net/http"
)

// JSON implements [Parser], it will parse request body into container
// Make sure your request body is not nil and is JSON-encoded or it will return an error
//
// Example:
//
//	var container = &struct{
//	    Name string `json:"name"`
//	    Age int `json:"age"`
//	    Tags []string `json:"tags"`
//	    VIP bool `json:"vip"`
//	}{}
//
//	err := JSON(request, container)
//	if err != nil {
//	    panic(err)
//	}
func JSON(request *http.Request, container any) error {
	if request == nil || request.Body == nil {
		return errors.New("invalid request")
	}
	decoder := json.NewDecoder(request.Body)
	return decoder.Decode(container)
}
