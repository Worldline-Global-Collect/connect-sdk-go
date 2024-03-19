// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// HostedFile represents class HostedFile
type HostedFile struct {
	FileName *string `json:"fileName,omitempty"`
	FileSize *string `json:"fileSize,omitempty"`
	FileType *string `json:"fileType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

// NewHostedFile constructs a new HostedFile
func NewHostedFile() *HostedFile {
	return &HostedFile{}
}
