package communicator

import (
	"net/url"
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.preprod.connect.worldline-solutions.com",
}

func TestToURIWithoutRequestParams(t *testing.T) {
	communicator := &Communicator{apiEndpoint: &baseURL}

	expectedURL, err := url.Parse("https://api.preprod.connect.worldline-solutions.com/v1/merchant/20000/convertamount")
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}

	uri1, err := communicator.toAbsoluteURI("v1/merchant/20000/convertamount", nil)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if uri1 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", uri1, *expectedURL)
	}

	uri2, err := communicator.toAbsoluteURI("/v1/merchant/20000/convertamount", nil)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if uri2 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", uri2, *expectedURL)
	}
}

func TestToURIWithRequestParams(t *testing.T) {
	amountParam, _ := communication.NewRequestParam("amount", "123")
	sourceParam, _ := communication.NewRequestParam("source", "USD")
	targetParam, _ := communication.NewRequestParam("target", "EUR")
	dummyParam, _ := communication.NewRequestParam("dummy", "é&%=")
	params := append(communication.RequestParams{}, *amountParam, *sourceParam, *targetParam, *dummyParam)

	communicator := &Communicator{apiEndpoint: &baseURL}

	expectedURL, err := url.Parse("https://api.preprod.connect.worldline-solutions.com/v1/merchant/20000/convertamount?amount=123&source=USD&target=EUR&dummy=%C3%A9%26%25%3D")
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}

	uri1, err := communicator.toAbsoluteURI("v1/merchant/20000/convertamount", params)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if uri1 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", uri1, *expectedURL)
	}

	uri2, err := communicator.toAbsoluteURI("/v1/merchant/20000/convertamount", params)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if uri2 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", uri2, *expectedURL)
	}
}
