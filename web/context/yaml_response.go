package context

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

type YAMLResponse struct {
	*Response
	data any
}

func (yamlResponse *YAMLResponse) SetContent(data any) IResponse {
	yamlResponse.data = data
	return yamlResponse
}

func (yamlResponse *YAMLResponse) Send(w http.ResponseWriter, r *http.Request) {
	yamlBytes, err := yaml.Marshal(yamlResponse.data)
	if err != nil {
		panic(err)
	}
	yamlResponse.content = yamlBytes
	yamlResponse.SetHeader("Content-Type", MIMEYAML)
	yamlResponse.Response.Send(w, r)
}
