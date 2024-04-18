// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// UploadDisputeFileResponse represents class UploadDisputeFileResponse
type UploadDisputeFileResponse struct {
	DisputeID *string `json:"disputeId,omitempty"`
	FileID    *string `json:"fileId,omitempty"`
}

// NewUploadDisputeFileResponse constructs a new UploadDisputeFileResponse instance
func NewUploadDisputeFileResponse() *UploadDisputeFileResponse {
	return &UploadDisputeFileResponse{}
}
