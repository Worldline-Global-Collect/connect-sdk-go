// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package apiv1

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
)

// Client represents a v1 client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Merchant represents the resource /{merchantId}
func (c *Client) Merchant(merchantID string) *merchant.Client {
	client, _ := merchant.NewClient(c.apiResource, map[string]string{
		"merchantId": merchantID,
	})
	return client
}

// NewClient constructs a v1 Client
//
// parent is the *communicator.APIResource on top of which we want to build the new v1 Client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
