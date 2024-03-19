package v1hmac

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
)

// Authenticator represents an authentication.Authenticator implementation using v1HMAC signatures
type Authenticator struct {
	apiKeyID     string
	secretAPIKey string
}

// GetAuthorization creates an authentication signature for the given httpMethod, resourceURI and requestHeaders
func (a Authenticator) GetAuthorization(httpMethod string, resourceURI url.URL, requestHeaders []communication.Header) (string, error) {
	if strings.TrimSpace(httpMethod) == "" {
		return "", errors.New("httpMethod is required")
	}
	dataToSign, err := toDataToSign(httpMethod, resourceURI, requestHeaders)
	if err != nil {
		return "", err
	}
	signedData, err := a.signData(dataToSign)
	if err != nil {
		return "", err
	}
	return "GCS v1HMAC:" + a.apiKeyID + ":" + signedData, nil
}

func (a Authenticator) signData(s string) (string, error) {
	mac := hmac.New(sha256.New, []byte(a.secretAPIKey))
	mac.Write([]byte(s))
	hmacOutput := mac.Sum(nil)

	writableBuffer := bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, &writableBuffer)
	_, err := encoder.Write(hmacOutput)
	if err != nil {
		return "", err
	}
	err = encoder.Close()
	if err != nil {
		return "", err
	}
	output := writableBuffer.String()
	return output, nil
}

func toDataToSign(httpMethod string, resourceURI url.URL, requestHeaders []communication.Header) (string, error) {
	xgcsHTTPHeaders := communication.Headers{}
	var contentType, date string
	for _, header := range requestHeaders {
		if strings.ToLower(header.Name()) == "content-type" {
			contentType = header.Value()
		} else if strings.ToLower(header.Name()) == "date" {
			date = header.Value()
		} else if strings.HasPrefix(strings.ToLower(header.Name()), "x-gcs") {
			newValue, err := toCanonicalizedHeaderValue(header.Value())
			if err != nil {
				return "", err
			}
			newHeader, err := communication.NewHeader(strings.ToLower(header.Name()), newValue)
			if err != nil {
				return "", err
			}

			xgcsHTTPHeaders = append(xgcsHTTPHeaders, *newHeader)
		}
	}
	sort.Sort(xgcsHTTPHeaders)

	dataToSign := bytes.Buffer{}
	dataToSign.WriteString(httpMethod)
	dataToSign.WriteRune('\n')
	dataToSign.WriteString(contentType)
	dataToSign.WriteRune('\n')
	dataToSign.WriteString(date)
	dataToSign.WriteRune('\n')
	for _, header := range xgcsHTTPHeaders {
		dataToSign.WriteString(header.Name() + ":" + header.Value())
		dataToSign.WriteRune('\n')
	}
	canonizalizedResource, err := toCanonicalizedResource(resourceURI)
	if err != nil {
		return "", err
	}
	dataToSign.WriteString(canonizalizedResource)
	dataToSign.WriteRune('\n')

	return dataToSign.String(), nil
}

func toCanonicalizedResource(resourceURI url.URL) (string, error) {
	resource := bytes.Buffer{}

	resource.WriteString(resourceURI.EscapedPath())
	if resourceURI.RawQuery != "" {
		query, err := url.QueryUnescape(resourceURI.RawQuery)
		if err != nil {
			return "", err
		}

		resource.WriteString("?" + query)
	}

	return resource.String(), nil
}

func toCanonicalizedHeaderValue(s string) (string, error) {
	regex, err := regexp.Compile("\r?\n[\t\f ]*")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(regex.ReplaceAllString(s, " ")), nil
}

// NewAuthenticator creates an Authenticator with the given apiKeyID and secretAPIKey
//- The apiKeyID is an identifier for the secret API key. The apiKeyID can be retrieved from the Configuration Center
//  This identifier is visible in the HTTP request and is also used to identify the correct account.
//- secretAPIKey is a shared secret. The shared secret can be retrieved from the Configuration Center.
//  An apiKeyID and secretAPIKey always go hand-in-hand, the difference is that secretAPIKey is never visible in the HTTP request.
//  This secret is used as input for the HMAC algorithm.
func NewAuthenticator(apiKeyID string, secretAPIKey string) (*Authenticator, error) {
	if strings.TrimSpace(apiKeyID) == "" {
		return nil, errors.New("apiKeyID is required")
	}
	if strings.TrimSpace(secretAPIKey) == "" {
		return nil, errors.New("secretAPIKey is required")
	}

	return &Authenticator{apiKeyID: apiKeyID, secretAPIKey: secretAPIKey}, nil
}
