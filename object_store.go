package pb

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

type object interface {
	Hash() string
	String() string
}

func putObject(o object) error {
	filename := fmt.Sprintf("./.pb/objects/%s", o.Hash())
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(o.String()))
	f.WriteString(encoded)
	f.Close()

	return nil
}

func getObject(k string) (string, error) {
	filename := fmt.Sprintf("./.pb/objects/%s", k)
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	if err != nil {
		return "", err
	}
	f.Close()

	decoded, err := base64.StdEncoding.DecodeString(string(buf.Bytes()))
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
