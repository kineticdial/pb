package pb

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

type Blob struct {
	contents string
}

func NewBlob(contents string) *Blob {
	return &Blob{contents}
}

func GetBlob(k string) (*Blob, error) {
	filename := fmt.Sprintf("./.pb/objects/%s", k)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}
	f.Close()

	decoded, err := base64.StdEncoding.DecodeString(string(buf.Bytes()))
	if err != nil {
		return nil, err
	}

	b := NewBlob(string(decoded))
	return b, nil
}

func (b *Blob) Put() error {
	filename := fmt.Sprintf("./.pb/objects/%s", b.Hash())
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(b.ToString()))
	f.WriteString(encoded)
	f.Close()

	return nil
}

func (b *Blob) ToString() string {
	return b.contents
}

func (b *Blob) Hash() string {
	h := sha1.New()
	h.Write([]byte(b.contents))
	return fmt.Sprintf("%x", h.Sum(nil))
}
