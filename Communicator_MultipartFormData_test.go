package connectsdk

import (
	"io"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/Worldline-Global-Collect/connect-sdk-go/authentication/v1hmac"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator"
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
	"github.com/Worldline-Global-Collect/connect-sdk-go/configuration"
	"github.com/Worldline-Global-Collect/connect-sdk-go/domain"
	"github.com/Worldline-Global-Collect/connect-sdk-go/json"
)

func TestMultipartFormDataUploadPostMultipartFormDataObjectWithResponse(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = comm.Post("/post", nil, nil, *multipart, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPostMultipartFormDataObjectPointerWithResponse(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = comm.Post("/post", nil, nil, multipart, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPostMultipartFormDataRequestWithResponse(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = comm.Post("/post", nil, nil, &MultipartFormDataObjectWrapper{multipart}, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPostMultipartFormDataObjectWithBodyHandler(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	err = comm.PostWithHandler("/post", nil, nil, *multipart, nil, func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller := json.DefaultMarshaller()
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartFormDataUploadPostMultipartFormDataObjectPointerWithBodyHandler(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	err = comm.PostWithHandler("/post", nil, nil, multipart, nil, func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller := json.DefaultMarshaller()
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartFormDataUploadPostMultipartFormDataRequestWithBodyHandler(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	err = comm.PostWithHandler("/post", nil, nil, &MultipartFormDataObjectWrapper{multipart}, nil, func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller := json.DefaultMarshaller()
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataObjectWithResponse(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = comm.Put("/put", nil, nil, *multipart, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataObjectPointerWithResponse(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = comm.Put("/put", nil, nil, multipart, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataRequestWithResponse(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = comm.Put("/put", nil, nil, &MultipartFormDataObjectWrapper{multipart}, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataObjectWithBodyHandler(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	err = comm.PutWithHandler("/put", nil, nil, *multipart, nil, func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller := json.DefaultMarshaller()
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataObjectPointerWithBodyHandler(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	err = comm.PutWithHandler("/put", nil, nil, multipart, nil, func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller := json.DefaultMarshaller()
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataRequestWithBodyHandler(t *testing.T) {
	comm, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer func(comm *communicator.Communicator) {
		_ = comm.Close()
	}(comm)

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := domain.NewUploadableFile("file.txt", content, "text/plain")
	_ = multipart.AddFile("file", *file)
	_ = multipart.AddValue("value", "Hello World")

	err = comm.PutWithHandler("/put", nil, nil, &MultipartFormDataObjectWrapper{multipart}, nil, func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller := json.DefaultMarshaller()
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func createCommunicator() (*communicator.Communicator, error) {
	httpBinURL := os.Getenv("httpbin.url")
	if httpBinURL == "" {
		httpBinURL = "http://httpbin.org"
	}
	apiEndpoint, err := url.Parse(httpBinURL)
	if err != nil {
		return nil, err
	}

	config, _ := configuration.DefaultV1HMACConfiguration("dummy", "dummy", "Worldline")
	config.APIEndpoint = *apiEndpoint

	connection, err := communicator.NewDefaultConnection(config.SocketTimeout,
		config.ConnectTimeout,
		config.KeepAliveTimeout,
		config.IdleTimeout,
		config.MaxConnections,
		config.Proxy)
	if err != nil {
		return nil, err
	}

	metadataProviderBuilder, _ := communicator.NewMetadataProviderBuilder(config.Integrator)
	metadataProviderBuilder.ShoppingCartExtension = config.ShoppingCartExtension

	metadataProvider, err := metadataProviderBuilder.Build()
	if err != nil {
		return nil, err
	}

	authenticator, err := v1hmac.NewAuthenticator(config.GetAPIKeyID(), config.GetSecretAPIKey())
	if err != nil {
		return nil, err
	}

	marshaller := json.DefaultMarshaller()

	return communicator.NewCommunicator(&config.APIEndpoint, connection, authenticator, metadataProvider, marshaller)
}

type HTTPBinResponse struct {
	Form  map[string]string `json:"form"`
	Files map[string]string `json:"files"`
}

type MultipartFormDataObjectWrapper struct {
	Multipart *communication.MultipartFormDataObject
}

func (w *MultipartFormDataObjectWrapper) ToMultipartFormDataObject() *communication.MultipartFormDataObject {
	return w.Multipart
}
