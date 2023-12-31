package context

import (
	"net/http"

	"github.com/wardonne/gopi/support/utils"
)

// JSONResponse used to response json format data
type JSONResponse struct {
	*Response
	data any
}

// SetContent sets response body content
func (jsonResponse *JSONResponse) SetContent(data any) {
	jsonResponse.data = data
}

// Send sends the response
func (jsonResponse *JSONResponse) Send(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := utils.JSONEncode(jsonResponse.data)
	if err != nil {
		panic(err)
	}
	jsonResponse.content = jsonBytes
	jsonResponse.SetHeader("Content-Type", MIMEJSON)
	jsonResponse.Response.Send(w, r)
}
