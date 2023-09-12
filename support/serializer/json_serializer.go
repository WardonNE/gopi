package serializer

import "encoding/json"

type JSONSerializer interface {
	json.Marshaler
	json.Unmarshaler
}
