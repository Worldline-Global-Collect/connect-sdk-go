// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DisplayedData represents class DisplayedData
type DisplayedData struct {
	DisplayedDataType *string         `json:"displayedDataType,omitempty"`
	RenderingData     *string         `json:"renderingData,omitempty"`
	ShowData          *[]KeyValuePair `json:"showData,omitempty"`
}

// NewDisplayedData constructs a new DisplayedData instance
func NewDisplayedData() *DisplayedData {
	return &DisplayedData{}
}
