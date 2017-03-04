package lib

import "crypto/sha256"
import "fmt"

// Blob represents the text inside a file that a TreeRef may point to.
type Blob struct {
	Contents string // Contents of file
}

// GetBlob fetches a Blob from the file k/v store by its SHA256 hash.
func GetBlob(k string) (*Blob, error) {
	contents, err := getObject(k)
	if err != nil {
		return nil, err
	}

	b := &Blob{contents}
	return b, nil
}

// Put stores a Blob in memory into the file k/v store.
func (b *Blob) Put() error {
	return putObject(b)
}

// String returns a Blob's contents.
func (b *Blob) String() string {
	return b.Contents
}

// Hash calculates a SHA256 hash of the Blob's contents.
func (b *Blob) Hash() string {
	h := sha256.New()
	h.Write([]byte(b.Contents))
	return fmt.Sprintf("%x", h.Sum(nil))
}
