// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package disputes

import (
	"github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication"
	"github.com/Worldline-Global-Collect/connect-sdk-go/domain"
)

// UploadFileRequest represents multipart/form-data parameters for Upload File
// Documentation can be found at https://apireference.connect.worldline-solutions.com/fileserviceapi/v1/en_US/go/disputes/uploadFile.html
type UploadFileRequest struct {
	File *domain.UploadableFile
}

// ToMultipartFormDataObject converts the multipart/form-data request to communication.MultipartFormDataObject
func (request UploadFileRequest) ToMultipartFormDataObject() *communication.MultipartFormDataObject {
	multipart, _ := communication.NewMultipartFormDataObject()

	if request.File != nil {
		_ = multipart.AddFile("file", *request.File)
	}

	return multipart
}

// NewUploadFileRequest constructs an instance of UploadFileRequest
func NewUploadFileRequest() *UploadFileRequest {
	return &UploadFileRequest{}
}
