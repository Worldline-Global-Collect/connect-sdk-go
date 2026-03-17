package json

import "io"

// Marshaller is the interface used to marshal and unmarshal Worldline Global Collect platform request
// and response objects to and from JSON. Thread-safe.
type Marshaller interface {
	// Marshal converts a request object to a JSON string.
	Marshal(v interface{}) (string, error)

	// Unmarshal converts a JSON string to a response object.
	Unmarshal(data string, v interface{}) error

	// UnmarshalFromReader converts the contents from an io.Reader to a response object.
	UnmarshalFromReader(reader io.Reader, v interface{}) error
}
