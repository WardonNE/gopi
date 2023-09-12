package context

import (
	"encoding/xml"
	"net/http"
)

type XMLResponse struct {
	*Response
	data any
}

func (xmlResponse *XMLResponse) SetContent(data any) IResponse {
	xmlResponse.data = data
	return xmlResponse
}

func (xmlResponse *XMLResponse) Send(w http.ResponseWriter, r *http.Request) {
	xmlBytes, err := xml.Marshal(xmlResponse.data)
	if err != nil {
		panic(err)
	}
	xmlResponse.content = xmlBytes
	xmlResponse.SetHeader("Content-Type", MIMEXML)
	xmlResponse.Response.Send(w, r)
}
