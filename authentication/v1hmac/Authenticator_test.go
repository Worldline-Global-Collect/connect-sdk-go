package v1hmac

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

func TestToCanonicalizedHeaderValue(t *testing.T) {
	result, err := toCanonicalizedHeaderValue("foo\nbar  ")
	if err != nil {
		t.Fatalf("TestToCanonicalizedHeaderValue : %v", err)
	}
	if result != "foo bar" {
		t.Fatalf("TestToCanonicalizedHeaderValue : wrong canonization 1 '%s' != '%s'",
			result, "foo bar")
	}

	result, err = toCanonicalizedHeaderValue("foo\r\n  bar")
	if err != nil {
		t.Fatalf("TestToCanonicalizedHeaderValue : %v", err)
	}
	if result != "foo bar" {
		t.Fatalf("TestToCanonicalizedHeaderValue : wrong canonization 2 '%s' != '%s'",
			result, "foo bar")
	}

	result, err = toCanonicalizedHeaderValue(" some value  \r\n \n with  some \r\n  spaces ")
	if err != nil {
		t.Fatalf("TestToCanonicalizedHeaderValue : %v", err)
	}
	if result != "some value    with  some  spaces" {
		t.Fatalf("TestToCanonicalizedHeaderValue : wrong canonization 3 '%s' != '%s'",
			result, "some value    with  some  spaces")
	}
}

func TestToDataToSign(t *testing.T) {
	var httpHeaders []communication.Header
	serverMetaInfoHeader, err := communication.NewHeader("X-GCS-ServerMetaInfo", "{\"platformIdentifier\":\"Windows 7/6.1 Java/1.7 (Oracle Corporation; Java HotSpot(TM) 64-Bit Server VM; 1.7.0_45)\",\"sdkIdentifier\":\"1.0\"}")
	if err != nil {
		t.Fatalf("TestToDataToSign : %v", err)
	}
	httpHeaders = append(httpHeaders, *serverMetaInfoHeader)

	contentTypeHeader, _ := communication.NewHeader("Content-Type", "application/json")
	httpHeaders = append(httpHeaders, *contentTypeHeader)

	clientMetaInfoHeader, _ := communication.NewHeader("X-GCS-ClientMetaInfo", "{\"foo\",\"bar\"}")
	httpHeaders = append(httpHeaders, *clientMetaInfoHeader)

	userAgentHeader, _ := communication.NewHeader("User-Agent", "Apache-HttpClient/4.3.4 (java 1.5)")
	httpHeaders = append(httpHeaders, *userAgentHeader)

	dateHeader, _ := communication.NewHeader("Date", "Mon, 07 Jul 2014 12:12:40 GMT")
	httpHeaders = append(httpHeaders, *dateHeader)

	urlArg, err := url.Parse("http://localhost:8080/v1/9991/services%20bla/convert/amount?foo=bar&mies=no%20bar")
	if err != nil {
		t.Fatalf("TestToDataToSign : %v", err)
	}

	dataToSign, err := toDataToSign("POST", *urlArg, httpHeaders)
	if err != nil {
		t.Fatalf("TestToDataToSign : %v", err)
	}

	expectedStart := "POST\napplication/json\n"
	expectedEnd := "x-gcs-clientmetainfo:{\"foo\",\"bar\"}\nx-gcs-servermetainfo:{\"platformIdentifier\":\"Windows 7/6.1 Java/1.7 (Oracle Corporation; Java HotSpot(TM) 64-Bit Server VM; 1.7.0_45)\",\"sdkIdentifier\":\"1.0\"}\n/v1/9991/services%20bla/convert/amount?foo=bar&mies=no bar\n"

	if !strings.HasPrefix(dataToSign, expectedStart) {
		t.Fatalf("TestToDataToSign : '%s' does not start with '%s'", dataToSign, expectedStart)
	}
	if !strings.HasSuffix(dataToSign, expectedEnd) {
		t.Fatalf("TestToDataToSign : '%s' does not end with '%s'", dataToSign, expectedEnd)
	}
}

func TestCreateAuthenticationSignature(t *testing.T) {
	auth, err := NewAuthenticator("apiKeyId", "secretApiKey")
	if err != nil {
		t.Fatalf("TestCreateAuthenticationSignature : %v", err)
	}

	dataToSign := "DELETE\napplication/json\nFri, 06 Jun 2014 13:39:43 GMT\nx-gcs-clientmetainfo:processed header value\nx-gcs-customerheader:processed header value\nx-gcs-servermetainfo:processed header value\n/v1/9991/tokens/123456789\n"

	authenticationSignature, err := auth.signData(dataToSign)
	if err != nil {
		t.Fatalf("TestCreateAuthenticationSignature : %v", err)
	}
	if authenticationSignature != "VfnXpPBQQoHZivTcAg0JvOWkhnzlPnaCPKpTQn/uMJM=" {
		t.Fatalf("TestCreateAuthenticationSignature : authenticationSignature '%s' != '%s'",
			"VfnXpPBQQoHZivTcAg0JvOWkhnzlPnaCPKpTQn/uMJM=", authenticationSignature)
	}
}

func TestCreateAuthenticationSignature2(t *testing.T) {
	auth, err := NewAuthenticator("EC36A74A98D21", "6Kj5HT0MQKC6D8eb7W3lTg71kVKVDSt1")
	if err != nil {
		t.Fatalf("TestCreateAuthenticationSignature2 : %v", err)
	}

	dataToSign := "GET\n\nFri, 06 Jun 2014 13:39:43 GMT\n/v1/9991/tokens/123456789\n"

	authenticationSignature, err := auth.signData(dataToSign)
	if err != nil {
		t.Fatalf("TestCreateAuthenticationSignature2 : %v", err)
	}
	if authenticationSignature != "9ond5EIN05dBXJGCLRK5om9pxHsyrh/12pZJ7bvmwNM=" {
		t.Fatalf("TestCreateAuthenticationSignature2 : authenticationSignature '%s' != '%s'",
			"9ond5EIN05dBXJGCLRK5om9pxHsyrh/12pZJ7bvmwNM=", authenticationSignature)
	}
}

func TestTotalMinimalExample(t *testing.T) {
	authenticator, err := NewAuthenticator("5e45c937b9db33ae", "I42Zf4pVnRdroHfuHnRiJjJ2B6+22h0yQt/R3nZR8Xg=")
	if err != nil {
		t.Fatalf("TestTotalMinimalExample : %v", err)
	}

	var httpHeaders []communication.Header

	userAgentHeader, _ := communication.NewHeader("User-Agent", "Apache-HttpClient/4.3.4 (java 1.5)")
	httpHeaders = append(httpHeaders, *userAgentHeader)

	dateHeader, _ := communication.NewHeader("Date", "Fri, 06 Jun 2014 13:39:43 GMT")
	httpHeaders = append(httpHeaders, *dateHeader)

	parsedURL, err := url.Parse("http://api.connect.worldline-solutions.com:8080/v1/9991/tokens/123456789")
	if err != nil {
		t.Fatalf("TestTotalMinimalExample : %v", err)
	}

	authorization, err := authenticator.GetAuthorization(http.MethodGet, *parsedURL, httpHeaders)
	if err != nil {
		t.Fatalf("TestTotalMinimalExample : %v", err)
	}
	if authorization != "GCS v1HMAC:5e45c937b9db33ae:J5LjfSBvrQNhu7gG0gvifZt+IWNDReGCmHmBmth6ueI=" {
		t.Fatalf("TestTotalMinimalExample : authorization '%s' != '%s'",
			"GCS v1HMAC:5e45c937b9db33ae:J5LjfSBvrQNhu7gG0gvifZt+IWNDReGCmHmBmth6ueI=", authorization)
	}
}

func TestTotalFullExample(t *testing.T) {
	authenticator, err := NewAuthenticator("5e45c937b9db33ae", "I42Zf4pVnRdroHfuHnRiJjJ2B6+22h0yQt/R3nZR8Xg=")
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}

	var httpHeaders []communication.Header

	userAgentHeader, _ := communication.NewHeader("User-Agent", "Apache-HttpClient/4.3.4 (java 1.5)")
	httpHeaders = append(httpHeaders, *userAgentHeader)

	serverMetaInfoHeader, _ := communication.NewHeader("X-GCS-ServerMetaInfo", "processed header value")
	httpHeaders = append(httpHeaders, *serverMetaInfoHeader)

	clientMetaInfoHeader, _ := communication.NewHeader("X-GCS-ClientMetaInfo", "processed header value")
	httpHeaders = append(httpHeaders, *clientMetaInfoHeader)

	contentTypeHeader, _ := communication.NewHeader("Content-Type", "application/json")
	httpHeaders = append(httpHeaders, *contentTypeHeader)

	customerHeader, _ := communication.NewHeader("X-GCS-CustomerHeader", "processed header value")
	httpHeaders = append(httpHeaders, *customerHeader)

	dateHeader, _ := communication.NewHeader("Date", "Fri, 06 Jun 2014 13:39:43 GMT")
	httpHeaders = append(httpHeaders, *dateHeader)

	parsedURL, err := url.Parse("http://api.connect.worldline-solutions.com:8080/v1/9991/tokens/123456789")
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}

	authorization, err := authenticator.GetAuthorization(http.MethodDelete, *parsedURL, httpHeaders)
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}
	if authorization != "GCS v1HMAC:5e45c937b9db33ae:jGWLz3ouN4klE+SkqO5gO+KkbQNM06Rric7E3dcfmqw=" {
		t.Fatalf("TestTotalFullExample : authorization '%s' != '%s'",
			"GCS v1HMAC:5e45c937b9db33ae:jGWLz3ouN4klE+SkqO5gO+KkbQNM06Rric7E3dcfmqw=", authorization)
	}
}
