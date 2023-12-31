package utils

import (
	"encoding/json"

	"github.com/wardonne/gopi/support"

	jsoniter "github.com/json-iterator/go"
)

// JSONEncode encodes data to json format
func JSONEncode(data any) (jsonBytes []byte, err error) {
	switch v := data.(type) {
	case json.Marshaler:
		jsonBytes, err = json.Marshal(v)
	case support.Mapable:
		jsonBytes, err = jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(v.ToMap())
	case support.Arrayable:
		jsonBytes, err = json.Marshal(v.ToArray())
	default:
		jsonBytes, err = json.Marshal(v)
	}
	return
}

// JSONDecode decodes json-encoded data to container
func JSONDecode(data []byte, container any) error {
	return json.Unmarshal(data, container)
}
