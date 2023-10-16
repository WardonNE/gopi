package binding

import (
	"errors"
	"net/http"

	"gopkg.in/yaml.v3"
)

// YAML implements [Parser], it parses request body into container
// Make sure your request body is not nil and is YAML-encoded, or it will returns an error
//
// Example:
//
//	var container = &struct{
//	    Name string `yaml:"name"`
//	    Age int `yaml:"age"`
//	    Tags []string `yaml:"tags"`
//	    VIP bool `yaml:"vip"`
//	}{}
//
//	err := YAML(request, container)
//	if err != nil {
//	    panic(err)
//	}
func YAML(request *http.Request, container any) error {
	if request == nil || request.Body == nil {
		return errors.New("invalid request")
	}
	decoder := yaml.NewDecoder(request.Body)
	return decoder.Decode(container)
}
