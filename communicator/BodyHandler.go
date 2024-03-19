package communicator

import (
	"io"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

// BodyHandler is a function for handling incoming body streams
type BodyHandler func(headers []communication.Header, reader io.Reader) error
