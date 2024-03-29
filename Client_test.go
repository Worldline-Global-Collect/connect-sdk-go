package connectsdk

import (
	"encoding/base64"
	"net/url"
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/configuration"
	"github.com/Worldline-Global-Collect/connect-sdk-go/json"
)

var testConfiguration = configuration.CommunicatorConfiguration{
	APIEndpoint: url.URL{
		Scheme: "https",
		Host:   "api.preprod.connect.worldline-solutions.com",
	},
	AuthorizationType:   configuration.V1HMAC,
	MaxConnections:      100,
	AuthorizationID:     "someKey",
	AuthorizationSecret: "someSecret",
	Integrator:          "Worldline",
}

func TestWithClientMetaInfo(t *testing.T) {
	client1, err := CreateClientFromConfiguration(&testConfiguration)
	if err != nil {
		t.Fatalf("TestWithClientMetaInfo: %v", err)
	}
	if client1.apiResource == nil {
		t.Fatal("TestWithClientMetaInfo: nil apiResource")
	}
	headers := client1.apiResource.ClientHeaders()
	if headers != nil {
		t.Fatal("TestWithClientMetaInfo: headers not nil")
	}

	client2, err := client1.WithClientMetaInfo("")
	if err != nil {
		t.Fatalf("TestWithClientMetaInfo: %v", err)
	}
	if client1 != client2 {
		t.Fatal("TestWithClientMetaInfo: clients not equal")
	}

	marshaller := json.DefaultMarshaller()
	clientMetaInfo, err := marshaller.Marshal(map[string]string{
		"test": "test",
	})
	if err != nil {
		t.Fatalf("TestWithClientMetaInfo: %v", err)
	}

	client3, err := client1.WithClientMetaInfo(clientMetaInfo)
	if err != nil {
		t.Fatalf("TestWithClientMetaInfo: %v", err)
	}
	if client3 == client1 {
		t.Fatal("TestWithClientMetaInfo: clients equal")
	}
	headers = client3.apiResource.ClientHeaders()
	if len(headers) != 1 {
		t.Fatalf("TestWithClientMetaInfo: invalid headers len %d", len(headers))
	}
	headerValue := base64.StdEncoding.EncodeToString([]byte(clientMetaInfo))
	if headers[0].Value() != headerValue {
		t.Fatalf("TestWithClientMetaInfo: header value mismatch '%v' '%v'", headers[0].Value(), headerValue)
	}

	client4, err := client3.WithClientMetaInfo(clientMetaInfo)
	if err != nil {
		t.Fatalf("TestWithClientMetaInfo: %v", err)
	}
	if client3 != client4 {
		t.Fatal("TestWithClientMetaInfo: clients not equal")
	}

	client5, err := client3.WithClientMetaInfo("")
	if err != nil {
		t.Fatal("TestWithClientMetaInfo: clients not equal")
	}
	if client3 == client5 {
		t.Fatal("TestWithClientMetaInfo: clients equal")
	}
}
