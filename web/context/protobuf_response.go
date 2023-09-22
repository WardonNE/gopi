package context

import (
	"net/http"

	"google.golang.org/protobuf/proto"
)

type ProtobufResponse struct {
	*Response
	data any
}

func (protobufResponse *ProtobufResponse) SetContent(data any) IResponse {
	protobufResponse.data = data
	return protobufResponse
}

func (protobufResponse *ProtobufResponse) Send(w http.ResponseWriter, r *http.Request) {
	protobufBytes, err := proto.Marshal(protobufResponse.data.(proto.Message))
	if err != nil {
		panic(err)
	}
	protobufResponse.content = protobufBytes
	protobufResponse.SetHeader("Content-Type", MIMEPROTOBUF)
}
