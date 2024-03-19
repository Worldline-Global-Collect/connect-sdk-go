package communicator

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
	"time"

	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
	commErrors "github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors"
	"github.com/Worldline-Global-Collect/connect-sdk-go/logging"
	"github.com/Worldline-Global-Collect/connect-sdk-go/logging/obfuscation"
)

// DefaultConnection is the default implementation for the connection interface. Supports Pooling, and is thread safe.
type DefaultConnection struct {
	client              http.Client
	underlyingTransport *http.Transport
	logger              logging.CommunicatorLogger
	proxyAuth           string
	bodyObfuscator      obfuscation.BodyObfuscator
	headerObfuscator    obfuscation.HeaderObfuscator
}

// Get sends a GET request to the Worldline Global Collect platform and calls the given response handler with the response.
func (c *DefaultConnection) Get(uri url.URL, headerList []communication.Header, handler communication.ResponseHandler) (interface{}, error) {
	return c.sendRequest("GET", uri, headerList, nil, "", handler)
}

// Delete sends a DELETE request to the Worldline Global Collect platform and calls the given response handler with the response.
func (c *DefaultConnection) Delete(uri url.URL, headerList []communication.Header, handler communication.ResponseHandler) (interface{}, error) {
	return c.sendRequest("DELETE", uri, headerList, nil, "", handler)
}

// Post sends a POST request to the Worldline Global Collect platform and calls the given response handler with the response.
func (c *DefaultConnection) Post(uri url.URL, headerList []communication.Header, body string, handler communication.ResponseHandler) (interface{}, error) {
	return c.sendRequest("POST", uri, headerList, strings.NewReader(body), body, handler)
}

// PostMultipart sends a multipart/form-data POST request to the Worldline Global Collect platform and calls the given response handler with the response.
func (c *DefaultConnection) PostMultipart(uri url.URL, headerList []communication.Header, body *communication.MultipartFormDataObject, handler communication.ResponseHandler) (interface{}, error) {
	r, err := c.createMultipartReader(body)
	if err != nil {
		return nil, err
	}
	defer func(r io.ReadCloser) {
		_ = r.Close()
	}(r)
	return c.sendRequest("POST", uri, headerList, r, "<binary content>", handler)
}

// Put sends a PUT request to the Worldline Global Collect platform and returns the response.
func (c *DefaultConnection) Put(uri url.URL, headerList []communication.Header, body string, handler communication.ResponseHandler) (interface{}, error) {
	return c.sendRequest("PUT", uri, headerList, strings.NewReader(body), body, handler)
}

// PutMultipart sends a multipart/form-data POST request to the Worldline Global Collect platform and calls the given response handler with the response.
func (c *DefaultConnection) PutMultipart(uri url.URL, headerList []communication.Header, body *communication.MultipartFormDataObject, handler communication.ResponseHandler) (interface{}, error) {
	r, err := c.createMultipartReader(body)
	if err != nil {
		return nil, err
	}
	defer func(r io.ReadCloser) {
		_ = r.Close()
	}(r)
	return c.sendRequest("PUT", uri, headerList, r, "<binary content>", handler)
}

func (c *DefaultConnection) sendRequest(method string, uri url.URL, headerList []communication.Header, body io.Reader, bodyString string, handler communication.ResponseHandler) (interface{}, error) {
	id, err := communication.PseudoUUID()
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest(method, uri.String(), body)
	if err != nil {
		c.logError(id, err)
		return nil, err
	}

	for _, h := range headerList {
		httpRequest.Header[h.Name()] = append(httpRequest.Header[h.Name()], h.Value())
	}
	if len(c.proxyAuth) > 0 {
		httpRequest.Header["Proxy-Authorization"] = append(httpRequest.Header["Proxy-Authorization"], c.proxyAuth)
	}

	start := time.Now()

	c.logRequest(id, bodyString, httpRequest)

	httpResponse, err := c.client.Do(httpRequest)
	switch ce := err.(type) {
	case *url.Error:
		{
			c.logError(id, ce)

			newErr, _ := commErrors.NewCommunicationError(ce)
			return nil, newErr
		}
	}
	if err != nil {
		c.logError(id, err)
		return nil, err
	}

	end := time.Now()

	defer func() {
		err := httpResponse.Body.Close()
		if err != nil {
			c.logError(id, err)
		}
	}()

	var responseHeaders []communication.Header
	for name, values := range httpResponse.Header {
		if name == "X-Gcs-Idempotence-Request-Timestamp" {
			name = "X-GCS-Idempotence-Request-Timestamp"
		}

		header, err := communication.NewHeader(name, values[0])
		if err != nil {
			c.logError(id, err)
			return nil, err
		}
		responseHeaders = append(responseHeaders, *header)
	}

	bodyReader := httpResponse.Body.(io.Reader)
	if isBinaryContent(responseHeaders) {
		c.logResponse(id, nil, true, httpResponse, end.Sub(start))
	} else {
		readBuffer := bytes.NewBuffer([]byte{})
		teeReader := io.TeeReader(httpResponse.Body, readBuffer)
		bodyReader = teeReader
		defer c.logResponse(id, io.MultiReader(readBuffer, teeReader), false, httpResponse, end.Sub(start))
	}

	return handler(httpResponse.StatusCode, responseHeaders, bodyReader)
}

func isBinaryContent(headers []communication.Header) bool {
	header := communication.Headers(headers).GetHeader("Content-Type")
	if header == nil {
		return false
	}

	contentType := strings.ToLower(header.Value())

	return !strings.HasPrefix(contentType, "text/") &&
		!strings.Contains(contentType, "json") &&
		!strings.Contains(contentType, "xml")
}

func (c *DefaultConnection) logRequest(id, body string, request *http.Request) {
	if c.logger == nil {
		return
	}

	var requestURL url.URL
	if request.URL != nil {
		requestURL = *request.URL
	}

	requestMessage, err := logging.NewRequestLogMessageBuilder(id, request.Method, requestURL, c.bodyObfuscator, c.headerObfuscator)
	if err != nil {
		c.logError(id, err)
		return
	}

	for k, v := range request.Header {
		for _, rv := range v {
			err = requestMessage.AddHeader(k, rv)
			if err != nil {
				c.logError(id, err)
				// continue with logging
			}
		}
	}

	err = requestMessage.SetBody(body, request.Header.Get("Content-Type"))
	if err != nil {
		c.logError(id, err)
		// continue with logging
	}

	message := requestMessage.BuildMessage()

	c.logger.Log(message.String())
}

func (c *DefaultConnection) logResponse(id string, reader io.Reader, binaryResponse bool, response *http.Response, duration time.Duration) {
	if c.logger == nil {
		return
	}

	responseMessage, err := logging.NewResponseLogMessageBuilder(id, response.StatusCode, duration, c.bodyObfuscator, c.headerObfuscator)
	if err != nil {
		c.logError(id, err)
		return
	}

	for k, v := range response.Header {
		for _, rv := range v {
			err = responseMessage.AddHeader(k, rv)
			if err != nil {
				c.logError(id, err)
				// continue with logging
			}
		}
	}

	if binaryResponse {
		err = responseMessage.SetBinaryBody(response.Header.Get("Content-Type"))
		if err != nil {
			c.logError(id, err)
			// continue with logging
		}
	} else {
		bodyBuff, err := ioutil.ReadAll(reader)
		if err != nil {
			c.logError(id, err)
			// continue with logging
		} else {
			err = responseMessage.SetBody(string(bodyBuff), response.Header.Get("Content-Type"))
			if err != nil {
				c.logError(id, err)
				// continue with logging
			}
		}

	}

	message := responseMessage.BuildMessage()

	c.logger.Log(message.String())
}

func (c *DefaultConnection) logError(id string, err error) {
	if c.logger != nil {
		c.logger.LogError(id, err)
	}
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func (c *DefaultConnection) createMultipartReader(body *communication.MultipartFormDataObject) (io.ReadCloser, error) {
	r, w := io.Pipe()

	writer := multipart.NewWriter(w)
	err := writer.SetBoundary(body.GetBoundary())
	if err != nil {
		return nil, err
	}
	if writer.FormDataContentType() != body.GetContentType() {
		return nil, errors.New("multipart.Writer  did not create the expected content type")
	}

	go func() {
		for name, value := range body.GetValues() {
			err := writer.WriteField(name, value)
			if err != nil {
				_ = w.CloseWithError(err)
				return
			}
		}
		for name, file := range body.GetFiles() {
			// Do not use writer.CreateFormFile because it does not allow a custom content type
			header := textproto.MIMEHeader{}
			header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(name), escapeQuotes(file.GetFileName())))
			header.Set("Content-Type", file.GetContentType())
			pw, err := writer.CreatePart(header)
			if err != nil {
				_ = w.CloseWithError(err)
				return
			}
			_, err = io.Copy(pw, file.GetContent())
			if err != nil {
				_ = w.CloseWithError(err)
				return
			}
		}
		err = writer.Close()
		if err != nil {
			_ = w.CloseWithError(err)
			return
		}
		_ = w.Close()
	}()

	return r, nil
}

// CloseIdleConnections closes all HTTP connections that have been idle for the specified time.
// This should also include all expired HTTP connections.
// timespan represents the time spent idle
// Note: in the current implementation, it is only possible to close the connection after a predetermined time
// Therefore, this implementation ignores the parameter, and instead uses the preconfigured one
func (c *DefaultConnection) CloseIdleConnections(t time.Duration) {
	// Assume t is equal to configured value
	c.underlyingTransport.CloseIdleConnections()
}

// CloseExpiredConnections closes all expired HTTP connections.
func (c *DefaultConnection) CloseExpiredConnections() {
	// No-op, because this is done automatically for this implementation
}

// SetBodyObfuscator sets the body obfuscator to use.
func (c *DefaultConnection) SetBodyObfuscator(bodyObfuscator obfuscation.BodyObfuscator) {
	c.bodyObfuscator = bodyObfuscator
}

// SetHeaderObfuscator sets the header obfuscator to use.
func (c *DefaultConnection) SetHeaderObfuscator(headerObfuscator obfuscation.HeaderObfuscator) {
	c.headerObfuscator = headerObfuscator
}

// EnableLogging implements the logging.Capable interface
// Enables logging to the given CommunicatorLogger
func (c *DefaultConnection) EnableLogging(logger logging.CommunicatorLogger) {
	c.logger = logger
}

// DisableLogging implements the logging.Capable interface
// Disables logging
func (c *DefaultConnection) DisableLogging() {
	c.logger = nil
}

// Close implements the io.Closer interface
func (c *DefaultConnection) Close() error {
	// No-op, because the http.Client connection's close automatically after the socket timeout passes
	// and they can't be closed manually
	return nil
}

// NewDefaultConnection creates a new object that implements Connection, and initializes it
func NewDefaultConnection(socketTimeout, connectTimeout, keepAliveTimeout, idleTimeout time.Duration, maxConnections int, proxy *url.URL) (*DefaultConnection, error) {
	dialer := net.Dialer{
		Timeout:   connectTimeout,
		KeepAlive: keepAliveTimeout,
	}

	transport := &http.Transport{
		DialContext:     dialer.DialContext,
		Proxy:           http.ProxyURL(proxy),
		IdleConnTimeout: idleTimeout,
		MaxIdleConns:    maxConnections,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	client := http.Client{
		Transport: transport,
		Timeout:   socketTimeout,
	}

	proxyAuth := ""
	if proxy != nil && proxy.User != nil {
		proxyAuth = "Basic " + base64.StdEncoding.EncodeToString([]byte(proxy.User.String()))
	}

	return &DefaultConnection{client, transport, nil, proxyAuth, obfuscation.DefaultBodyObfuscator(), obfuscation.DefaultHeaderObfuscator()}, nil
}
