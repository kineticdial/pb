package pb

import (
	"crypto/sha1"
	"fmt"
)

type Blob struct {
	Contents string
}

func GetBlob(k string) (*Blob, error) {
	contents, err := GetObject(k)
	if err != nil {
		return nil, err
	}

	b := &Blob{contents}
	return b, nil
}

func (b *Blob) Put() error {
	return PutObject(b)
}

func (b *Blob) String() string {
	return b.Contents
}

func (b *Blob) Hash() string {
	h := sha1.New()
	h.Write([]byte(b.Contents))
	return fmt.Sprintf("%x", h.Sum(nil))
}
