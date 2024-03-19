// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package merchant

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/captures"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/disputes"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/files"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/hostedcheckouts"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/hostedmandatemanagements"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/installments"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/mandates"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/payments"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/payouts"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/productgroups"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/products"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/refunds"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/riskassessments"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/services"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/sessions"
	"github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/tokens"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
)

// Client represents a merchant client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Hostedcheckouts represents the resource /{merchantId}/hostedcheckouts
func (c *Client) Hostedcheckouts() *hostedcheckouts.Client {
	client, _ := hostedcheckouts.NewClient(c.apiResource, nil)
	return client
}

// Hostedmandatemanagements represents the resource /{merchantId}/hostedmandatemanagements
func (c *Client) Hostedmandatemanagements() *hostedmandatemanagements.Client {
	client, _ := hostedmandatemanagements.NewClient(c.apiResource, nil)
	return client
}

// Payments represents the resource /{merchantId}/payments
func (c *Client) Payments() *payments.Client {
	client, _ := payments.NewClient(c.apiResource, nil)
	return client
}

// Captures represents the resource /{merchantId}/captures
func (c *Client) Captures() *captures.Client {
	client, _ := captures.NewClient(c.apiResource, nil)
	return client
}

// Refunds represents the resource /{merchantId}/refunds
func (c *Client) Refunds() *refunds.Client {
	client, _ := refunds.NewClient(c.apiResource, nil)
	return client
}

// Disputes represents the resource /{merchantId}/disputes
func (c *Client) Disputes() *disputes.Client {
	client, _ := disputes.NewClient(c.apiResource, nil)
	return client
}

// Payouts represents the resource /{merchantId}/payouts
func (c *Client) Payouts() *payouts.Client {
	client, _ := payouts.NewClient(c.apiResource, nil)
	return client
}

// Productgroups represents the resource /{merchantId}/productgroups
func (c *Client) Productgroups() *productgroups.Client {
	client, _ := productgroups.NewClient(c.apiResource, nil)
	return client
}

// Products represents the resource /{merchantId}/products
func (c *Client) Products() *products.Client {
	client, _ := products.NewClient(c.apiResource, nil)
	return client
}

// Riskassessments represents the resource /{merchantId}/riskassessments
func (c *Client) Riskassessments() *riskassessments.Client {
	client, _ := riskassessments.NewClient(c.apiResource, nil)
	return client
}

// Services represents the resource /{merchantId}/services
func (c *Client) Services() *services.Client {
	client, _ := services.NewClient(c.apiResource, nil)
	return client
}

// Tokens represents the resource /{merchantId}/tokens
func (c *Client) Tokens() *tokens.Client {
	client, _ := tokens.NewClient(c.apiResource, nil)
	return client
}

// Mandates represents the resource /{merchantId}/mandates
func (c *Client) Mandates() *mandates.Client {
	client, _ := mandates.NewClient(c.apiResource, nil)
	return client
}

// Sessions represents the resource /{merchantId}/sessions
func (c *Client) Sessions() *sessions.Client {
	client, _ := sessions.NewClient(c.apiResource, nil)
	return client
}

// Installments represents the resource /{merchantId}/installments
func (c *Client) Installments() *installments.Client {
	client, _ := installments.NewClient(c.apiResource, nil)
	return client
}

// Files represents the resource /{merchantId}/files
func (c *Client) Files() *files.Client {
	client, _ := files.NewClient(c.apiResource, nil)
	return client
}

// NewClient constructs a Merchant Client
//
// parent is the *communicator.APIResource on top of which we want to build the new Merchant Client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
