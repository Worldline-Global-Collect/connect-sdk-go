package connectsdk

import (
	"reflect"
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/authentication/v1hmac"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	"github.com/Worldline-Global-Collect/connect-sdk-go/configuration"
	"github.com/Worldline-Global-Collect/connect-sdk-go/json"
)

var apiKeyID = "someKey"
var secretAPIKey = "someSecret"

func TestCreateConfiguration(t *testing.T) {
	defaultConfiguration, _ := configuration.DefaultV1HMACConfiguration(apiKeyID, secretAPIKey, "Test")

	conf, err := CreateV1HMACConfiguration(apiKeyID, secretAPIKey, "Test")
	if err != nil {
		t.Fatalf("TestCreateConfiguration: %v", err)
	}

	if conf.APIEndpoint != defaultConfiguration.APIEndpoint {
		t.Fatalf("TestCreateConfiguration: APIEndpoint mismatch %v %v",
			conf.APIEndpoint, defaultConfiguration.APIEndpoint)
	}

	if conf.AuthorizationType != defaultConfiguration.AuthorizationType {
		t.Fatalf("TestCreateConfiguration: AuthorizationType mismatch %v %v",
			conf.AuthorizationType, defaultConfiguration.AuthorizationType)
	}

	if conf.IdleTimeout != defaultConfiguration.IdleTimeout {
		t.Fatalf("TestCreateConfiguration: IdleTimeout mismatch %v %v",
			conf.IdleTimeout, defaultConfiguration.IdleTimeout)
	}

	if conf.ConnectTimeout != defaultConfiguration.ConnectTimeout {
		t.Fatalf("TestCreateConfiguration: ConnectTimeout mismatch %v %v",
			conf.ConnectTimeout, defaultConfiguration.ConnectTimeout)
	}

	if conf.SocketTimeout != defaultConfiguration.SocketTimeout {
		t.Fatalf("TestCreateConfiguration: SocketTimeout mismatch %v %v",
			conf.SocketTimeout, defaultConfiguration.SocketTimeout)
	}

	if conf.MaxConnections != defaultConfiguration.MaxConnections {
		t.Fatalf("TestCreateConfiguration: MaxConnections mismatch %v %v",
			conf.MaxConnections, defaultConfiguration.MaxConnections)
	}

	if conf.GetAPIKeyID() != apiKeyID {
		t.Fatalf("TestCreateConfiguration: APIKeyID mismatch %v %v",
			conf.GetAPIKeyID(), apiKeyID)
	}

	if conf.GetSecretAPIKey() != secretAPIKey {
		t.Fatalf("TestCreateConfiguration: SecretAPIKey mismatch %v %v",
			conf.GetSecretAPIKey(), secretAPIKey)
	}
}

func TestCreateCommunicator(t *testing.T) {
	marshaller := json.DefaultMarshaller()

	comm, err := CreateCommunicator(apiKeyID, secretAPIKey, "Test")
	if err != nil {
		t.Fatalf("TestCreateCommunicator: %v", err)
	}
	if comm.Marshaller() != marshaller {
		t.Fatalf("TestCreateCommunicator: marshaller mismatch %v %v",
			comm.Marshaller(), marshaller)
	}

	connection := comm.Connection()
	if _, isDefaultConnection := connection.(*communicator.DefaultConnection); !isDefaultConnection {
		t.Fatalf("TestCreateCommunicator: connection type mismatch %T", connection)
	}

	authenticator := comm.Authenticator()
	if _, isV1HMACAuthenticator := authenticator.(*v1hmac.Authenticator); !isV1HMACAuthenticator {
		t.Fatalf("TestCreateCommunicator: authenticator type mismatch %T", authenticator)
	}

	authAPIKeyID := reflect.Indirect(reflect.ValueOf(authenticator)).FieldByName("apiKeyID").String()
	if authAPIKeyID != apiKeyID {
		t.Fatalf("TestCreateCommunicator: apiKeyID mismatch %v", authAPIKeyID)
	}
	authSecretAPIKey := reflect.Indirect(reflect.ValueOf(authenticator)).FieldByName("secretAPIKey").String()
	if authSecretAPIKey != secretAPIKey {
		t.Fatalf("TestCreateCommunicator: secretAPIKey mismatch %v", authSecretAPIKey)
	}

	metadataProvider := comm.MetadataProvider()
	headers := metadataProvider.MetadataHeaders()
	if len(headers) != 1 {
		t.Fatalf("TestCreateCommunicator: headers len mismatch %v", len(headers))
	}
	if headers[0].Name() != "X-GCS-ServerMetaInfo" {
		t.Fatalf("TestCreateCommunicator: header name mismatch %v", headers[0].Name())
	}
}
