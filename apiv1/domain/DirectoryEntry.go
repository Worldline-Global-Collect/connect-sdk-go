// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// DirectoryEntry represents class DirectoryEntry
type DirectoryEntry struct {
	CountryNames *[]string `json:"countryNames,omitempty"`
	IssuerID     *string   `json:"issuerId,omitempty"`
	IssuerList   *string   `json:"issuerList,omitempty"`
	IssuerName   *string   `json:"issuerName,omitempty"`
}

// NewDirectoryEntry constructs a new DirectoryEntry instance
func NewDirectoryEntry() *DirectoryEntry {
	return &DirectoryEntry{}
}
