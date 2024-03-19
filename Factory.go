package connectsdk

import (
	"errors"
	"net/url"

	"github.com/Worldline-Global-Collect/connect-sdk-go/authentication"
	"github.com/Worldline-Global-Collect/connect-sdk-go/authentication/v1hmac"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	"github.com/Worldline-Global-Collect/connect-sdk-go/configuration"
	"github.com/Worldline-Global-Collect/connect-sdk-go/json"
)

// CreateV1HMACConfiguration creates a CommunicatorConfiguration with default v1HMAC settings and the given apiKeyID and secretAPIKey
func CreateV1HMACConfiguration(apiKeyID, secretAPIKey, integrator string) (*configuration.CommunicatorConfiguration, error) {
	return configuration.DefaultV1HMACConfiguration(apiKeyID, secretAPIKey, integrator)
}

// CreateCommunicatorBuilder creates a CommunicatorBuilder with the given CommunicatorConfiguration
func CreateCommunicatorBuilder(apiKeyID, secretAPIKey, integrator string) (*CommunicatorBuilder, error) {
	conf, err := CreateV1HMACConfiguration(apiKeyID, secretAPIKey, integrator)
	if err != nil {
		return nil, err
	}

	return CreateCommunicatorBuilderFromConfiguration(conf)
}

// CreateCommunicatorBuilderFromConfiguration creates a CommunicatorBuilder with the given CommunicatorConfiguration
func CreateCommunicatorBuilderFromConfiguration(config *configuration.CommunicatorConfiguration) (*CommunicatorBuilder, error) {
	builder := NewCommunicatorBuilder()

	connection, err := communicator.NewDefaultConnection(config.SocketTimeout,
		config.ConnectTimeout,
		config.KeepAliveTimeout,
		config.IdleTimeout,
		config.MaxConnections,
		config.Proxy)
	if err != nil {
		return nil, err
	}

	metadataProviderBuilder, err := communicator.NewMetadataProviderBuilder(config.Integrator)
	if err != nil {
		return nil, err
	}
	metadataProviderBuilder.ShoppingCartExtension = config.ShoppingCartExtension

	metadataProvider, err := metadataProviderBuilder.Build()
	if err != nil {
		return nil, err
	}

	authenticator, err := getAuthenticator(config)
	if err != nil {
		return nil, err
	}

	marshaller := json.DefaultMarshaller()

	return builder.
		WithAPIEndpoint(&config.APIEndpoint).
		WithConnection(connection).
		WithMetadataProvider(metadataProvider).
		WithAuthenticator(authenticator).
		WithMarshaller(marshaller), nil
}

func getAuthenticator(conf *configuration.CommunicatorConfiguration) (authentication.Authenticator, error) {
	if conf.AuthorizationType == configuration.V1HMAC {
		return v1hmac.NewAuthenticator(conf.GetAPIKeyID(), conf.GetSecretAPIKey())
	}
	return nil, errors.New("unknown authorizationType " + string(conf.AuthorizationType))
}

// CreateCommunicator creates a Communicator with default settings and the given apiKeyID and secretAPIKey
func CreateCommunicator(apiKeyID, secretAPIKey, integrator string) (*communicator.Communicator, error) {
	conf, err := CreateV1HMACConfiguration(apiKeyID, secretAPIKey, integrator)
	if err != nil {
		return nil, err
	}

	return CreateCommunicatorFromConfiguration(conf)
}

// CreateCommunicatorFromConfiguration creates a Communicator with the given CommunicatorConfiguration
func CreateCommunicatorFromConfiguration(conf *configuration.CommunicatorConfiguration) (*communicator.Communicator, error) {
	builder, err := CreateCommunicatorBuilderFromConfiguration(conf)
	if err != nil {
		return nil, err
	}

	return builder.Build()
}

// CreateCommunicatorWithDefaultMarshaller creates a Communicator with the given components and a default marshaller
func CreateCommunicatorWithDefaultMarshaller(apiEndpoint *url.URL, connection communicator.Connection, authenticator authentication.Authenticator, metadataProvider *communicator.MetadataProvider) (*communicator.Communicator, error) {
	return communicator.NewCommunicator(apiEndpoint, connection, authenticator, metadataProvider, json.DefaultMarshaller())
}

// CreateClient creates a Client with the given CommunicatorConfiguration
func CreateClient(apiKeyID, secretAPIKey, integrator string) (*Client, error) {
	conf, err := CreateV1HMACConfiguration(apiKeyID, secretAPIKey, integrator)
	if err != nil {
		return nil, err
	}

	return CreateClientFromConfiguration(conf)
}

// CreateClientFromConfiguration creates a Client with the given CommunicatorConfiguration
func CreateClientFromConfiguration(config *configuration.CommunicatorConfiguration) (*Client, error) {
	comm, err := CreateCommunicatorFromConfiguration(config)
	if err != nil {
		return nil, err
	}

	return CreateClientFromCommunicator(comm)
}

// CreateClientWithDefaultMarshaller creates a Client with the given components and a default marshaller
func CreateClientWithDefaultMarshaller(apiEndpoint *url.URL, connection communicator.Connection, authenticator authentication.Authenticator, metadataProvider *communicator.MetadataProvider) (*Client, error) {
	comm, err := CreateCommunicatorWithDefaultMarshaller(apiEndpoint, connection, authenticator, metadataProvider)
	if err != nil {
		return nil, err
	}

	return CreateClientFromCommunicator(comm)
}

// CreateClientFromCommunicator creates a Client with the given Communicator
func CreateClientFromCommunicator(communicator *communicator.Communicator) (*Client, error) {
	return NewClient(communicator)
}

// NewCallContext creates a CallContext using the given idempotenceKey
func NewCallContext(idempotenceKey string) *communicator.CallContext {
	return communicator.NewCallContext(idempotenceKey)
}
