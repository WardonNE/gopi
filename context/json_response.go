package context

import (
	"net/http"

	"github.com/wardonne/gopi/utils"
)

type JSONResponse struct {
	*Response
	data any
}

func (jsonResponse *JSONResponse) SetContent(data any) IResponse {
	jsonResponse.data = data
	return jsonResponse
}

func (jsonResponse *JSONResponse) Send(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := utils.JSONEncode(jsonResponse.data)
	if err != nil {
		panic(err)
	}
	jsonResponse.content = jsonBytes
	jsonResponse.SetHeader("Content-Type", MIMEJSON)
	jsonResponse.Response.Send(w, r)
}
