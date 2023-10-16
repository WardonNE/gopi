package context

import (
	"net/http"

	"google.golang.org/protobuf/proto"
)

// ProtobufResponse used to send a protobuf response
type ProtobufResponse struct {
	*Response
	data any
}

// SetContent sets response body content
//
// NOTICE: data should be [proto.Message]
func (protobufResponse *ProtobufResponse) SetContent(data any) IResponse {
	protobufResponse.data = data
	return protobufResponse
}

// Send sends the response
func (protobufResponse *ProtobufResponse) Send(w http.ResponseWriter, r *http.Request) {
	protobufBytes, err := proto.Marshal(protobufResponse.data.(proto.Message))
	if err != nil {
		panic(err)
	}
	protobufResponse.content = protobufBytes
	protobufResponse.SetHeader("Content-Type", MIMEPROTOBUF)
}
