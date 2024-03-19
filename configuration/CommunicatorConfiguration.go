package configuration

import (
	"net/url"
	"time"

	"github.com/Worldline-Global-Collect/connect-sdk-go/domain"
)

// CommunicatorConfiguration represents the configuration to be used by a Communicator
type CommunicatorConfiguration struct {
	// APIEndpoint represents the API endpoint of for the communicator
	// See https://apireference.connect.worldline-solutions.com/s2sapi/v1/en_US/go/endpoints.html
	APIEndpoint url.URL
	// ConnectTimeout represents the connect timeout
	ConnectTimeout time.Duration
	// SocketTimeout represents the request timeout
	SocketTimeout time.Duration
	// IdleTimeout represents the idle connection timeout
	IdleTimeout time.Duration
	// KeepAliveTimeout represents the HTTP KeepAlive interval
	KeepAliveTimeout time.Duration
	// MaxConnections represents the maximum amount of concurrent pooled connections
	MaxConnections int
	// AuthorizationType represents authorizationType used to sign the requests
	AuthorizationType AuthorizationType
	// AuthorizationID represents an id used for authorization. The meaning of this id is different for each authorization type.
	// For instance, for v1HMAC this is the identifier for the secret API key.
	AuthorizationID string
	// AuthorizationSecret represents a secret used for authorization. The meaning of this secret is different for each authorization type.
	// For instance, for v1HMAC this is the secret API key.
	AuthorizationSecret string
	// Proxy represents the URL for the connection proxy
	Proxy *url.URL
	// Integrator represents the integrator name
	Integrator string
	// ShoppingCartExtension represents the shopping cart extension used in the metadata headers
	ShoppingCartExtension *domain.ShoppingCartExtension
}

// GetAPIKeyID returns an identifier for the secret API key.
// The APIKeyID can be retrieved from the Configuration Center.
// This identifier is visible in the HTTP request and is also used to identify the correct account.
//
// This function is an alias for getting c.AuthorizationID
func (c *CommunicatorConfiguration) GetAPIKeyID() string {
	return c.AuthorizationID
}

// SetAPIKeyID sets an identifier for the secret API key.
// The APIKeyID can be retrieved from the Configuration Center.
// This identifier is visible in the HTTP request and is also used to identify the correct account.
//
// This function is an alias for setting c.AuthorizationID
func (c *CommunicatorConfiguration) SetAPIKeyID(apiKeyID string) {
	c.AuthorizationID = apiKeyID
}

// GetSecretAPIKey returns a shared secret.
// The shared secret can be retrieved from the Configuration Center.
// An APIKeyID and SecretAPIKey always go hand-in-hand, the difference is that SecretAPIKey is never visible in the HTTP request.
// This secret is used as input for the HMAC algorithm.
//
// This function is an alias for getting c.AuthorizationSecret
func (c *CommunicatorConfiguration) GetSecretAPIKey() string {
	return c.AuthorizationSecret
}

// SetSecretAPIKey sets a shared secret.
// The shared secret can be retrieved from the Configuration Center.
// An APIKeyID and SecretAPIKey always go hand-in-hand, the difference is that SecretAPIKey is never visible in the HTTP request.
// This secret is used as input for the HMAC algorithm.
//
// This function is an alias for setting c.AuthorizationSecret
func (c *CommunicatorConfiguration) SetSecretAPIKey(secretAPIKey string) {
	c.AuthorizationSecret = secretAPIKey
}

// The default configuration used by the factory is the following:
// APIEndpoint: api.connect.worldline-solutions.com
// ConnectTimeout: 5 seconds
// SocketTimeout: 30 seconds
// IdleTimeout: 5 seconds
// KeepAliveTimeout: 30 seconds
// MaxConnections: 10
// AuthorizationType: V1HMAC
// Proxy: none
var defaultConfiguration = CommunicatorConfiguration{
	APIEndpoint: url.URL{
		Scheme: "https",
		Host:   "api.connect.worldline-solutions.com",
	},
	ConnectTimeout:   5 * time.Second,
	SocketTimeout:    30 * time.Second,
	IdleTimeout:      5 * time.Second,
	KeepAliveTimeout: 30 * time.Second,
	MaxConnections:   10,
}

// DefaultV1HMACConfiguration returns the default communicator configuration for the v1HMAC authorization type
func DefaultV1HMACConfiguration(apiKeyID, secretAPIKey, integrator string) (*CommunicatorConfiguration, error) {
	customizedConfiguration := defaultConfiguration

	customizedConfiguration.Integrator = integrator
	customizedConfiguration.AuthorizationType = V1HMAC
	customizedConfiguration.AuthorizationID = apiKeyID
	customizedConfiguration.AuthorizationSecret = secretAPIKey

	return &customizedConfiguration, nil
}
