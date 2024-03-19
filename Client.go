// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package connectsdk

import (
	"encoding/base64"

	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	"github.com/Worldline-Global-Collect/connect-sdk-go/logging"
	"github.com/Worldline-Global-Collect/connect-sdk-go/logging/obfuscation"
)

// Client is the Worldline Global Collect platform client. Thread-safe.
//
// This client and all its child clients are bound to one specific value for the X-GCS-ClientMetaInfo header.
// To get a new client with a different header value, use WithClientMetaInfo.
type Client struct {
	apiResource *communicator.APIResource
}

// WithClientMetaInfo returns a new Client which uses the passed meta data for the X-GCS-ClientMetaInfo header.
// - clientMetaInfo is a JSON string containing the meta data for the client
// - can give an error if the given clientMetaInfo is not a valid JSON string
func (c *Client) WithClientMetaInfo(clientMetaInfo string) (*Client, error) {
	if len(c.apiResource.ClientMetaInfo()) == 0 && len(clientMetaInfo) == 0 {
		return c, nil
	}
	if len(clientMetaInfo) == 0 {
		return newClient(c.apiResource.Communicator(), "")
	}
	// Checking to see if this is valid JSON (no JSON parse errors)
	var testMap map[string]interface{}

	err := c.apiResource.Communicator().Marshaller().Unmarshal(clientMetaInfo, &testMap)
	if err != nil {
		return nil, err
	}

	clientMetaInfoBase64 := base64.StdEncoding.EncodeToString([]byte(clientMetaInfo))

	if clientMetaInfoBase64 == c.apiResource.ClientMetaInfo() {
		return c, nil
	}

	return newClient(c.apiResource.Communicator(), clientMetaInfoBase64)
}

// SetBodyObfuscator sets the body obfuscator to use.
func (c *Client) SetBodyObfuscator(bodyObfuscator obfuscation.BodyObfuscator) {
	c.apiResource.Communicator().SetBodyObfuscator(bodyObfuscator)
}

// SetHeaderObfuscator sets the header obfuscator to use.
func (c *Client) SetHeaderObfuscator(headerObfuscator obfuscation.HeaderObfuscator) {
	c.apiResource.Communicator().SetHeaderObfuscator(headerObfuscator)
}

// EnableLogging turns on logging using the given communicator logger.
func (c *Client) EnableLogging(communicatorLogger logging.CommunicatorLogger) {
	c.apiResource.Communicator().EnableLogging(communicatorLogger)
}

// DisableLogging turns off logging.
func (c *Client) DisableLogging() {
	c.apiResource.Communicator().DisableLogging()
}

// Close calls the internal closer of the communicator
func (c *Client) Close() error {
	return c.apiResource.Communicator().Close()
}

// V1 represents API v1
func (c *Client) V1() *apiv1.Client {
	client, _ := apiv1.NewClient(c.apiResource, nil)
	return client
}

// NewClient creates a new Client with the given communicator
func NewClient(communicator *communicator.Communicator) (*Client, error) {
	return newClient(communicator, "")
}

func newClient(comm *communicator.Communicator, clientMetaInfo string) (*Client, error) {
	apiResource, err := communicator.NewAPIResource(comm, clientMetaInfo, nil)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
