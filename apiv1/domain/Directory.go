// This class was auto-generated from the API references found at
// https://apireference.connect.worldline-solutions.com/

package domain

// Directory represents class Directory
type Directory struct {
	Entries *[]DirectoryEntry `json:"entries,omitempty"`
}

// NewDirectory constructs a new Directory instance
func NewDirectory() *Directory {
	return &Directory{}
}
