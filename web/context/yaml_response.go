package context

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

// YAMLResponse used to send YAML-encoded response
type YAMLResponse struct {
	*Response
	data any
}

// SetContent sets response body content
func (yamlResponse *YAMLResponse) SetContent(data any) IResponse {
	yamlResponse.data = data
	return yamlResponse
}

// Send sends the response
func (yamlResponse *YAMLResponse) Send(w http.ResponseWriter, r *http.Request) {
	yamlBytes, err := yaml.Marshal(yamlResponse.data)
	if err != nil {
		panic(err)
	}
	yamlResponse.content = yamlBytes
	yamlResponse.SetHeader("Content-Type", MIMEYAML)
	yamlResponse.Response.Send(w, r)
}
