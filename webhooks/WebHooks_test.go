package webhooks

import (
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/json"
)

func TestNewV1Helper(t *testing.T) {
	store := InMemorySecretKeyStore()

	helper, err := V1().NewHelper(store)
	if err != nil {
		t.Fatal(err)
	}

	marshaller := json.DefaultMarshaller()

	if helper.Marshaller() != marshaller {
		t.Fatalf("marshaller mismatch %v %v", helper.Marshaller(), marshaller)
	}
	if helper.SecretKeyStore() != store {
		t.Fatalf("secretKeyStore mismatch %v %v", helper.SecretKeyStore(), store)
	}
}
