package pb

import (
	"crypto/sha1"
	"fmt"
)

type Blob struct {
	Contents string
}

func NewBlob(contents string) *Blob {
	return &Blob{contents}
}

func GetBlob(k string) (*Blob, error) {
	var b Blob
	err := Get(k, &b)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (b *Blob) Put() error {
	err := Put(b.Hash(), b)
	if err != nil {
		return err
	}

	return nil
}

func (b *Blob) Hash() string {
	h := sha1.New()
	h.Write([]byte(b.Contents))
	return fmt.Sprintf("%x", h.Sum(nil))
}
