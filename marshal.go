package ds

import "encoding/json"

// SerializableJSON enables JSON marshaling and unmarshaling of a type.
type SerializableJSON interface {
	json.Marshaler
	json.Unmarshaler
}

// JSONMarshaler Encodes structure to JSON.
type JSONMarshaler func(v any) ([]byte, error)

// JSONUnmarshaler Decodes JSON data to structure.
type JSONUnmarshaler func(data []byte, v any) error

// DefaultJSONMarshaler global JSON marshaler function, uses std lib by default (encoding/json).
var DefaultJSONMarshaler JSONMarshaler = json.Marshal

// DefaultJSONUnmarshaler global JSON unmarshaler function, uses std lib by default (encoding/json).
var DefaultJSONUnmarshaler JSONUnmarshaler = json.Unmarshal
