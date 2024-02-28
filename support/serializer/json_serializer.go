package serializer

import "encoding/json"

// JSONSerializer json serializer
type JSONSerializer interface {
	json.Marshaler
	json.Unmarshaler
}
