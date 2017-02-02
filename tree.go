package pb

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Tree struct {
	Refs []*TreeRef
}

func GetTree(k string) (*Tree, error) {
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

	// Trim potential new line from end of file.
	decoded_string := strings.TrimRight(string(decoded), "\n")
	records := strings.Split(decoded_string, "\n")

	trs := make([]*TreeRef, 0, len(records))
	for _, r := range records {
		decoded_r, err := Decode(r)
		if err != nil {
			return nil, err
		}

		trs = append(trs, decoded_r)
	}

	t := &Tree{trs}
	return t, nil
}

func (t *Tree) Put() error {
	filename := fmt.Sprintf("./.pb/objects/%s", t.Hash())
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(t.String()))
	f.WriteString(encoded)
	f.Close()

	return nil
}

func (t *Tree) String() string {
	buf := bytes.NewBufferString("")

	sort.Sort(ByString(t.Refs))

	for _, tref := range t.Refs {
		buf.WriteString(tref.String())
		buf.WriteString("\n")
	}

	return buf.String()
}

func (t *Tree) Hash() string {
	h := sha1.New()
	h.Write([]byte(t.String()))

	return fmt.Sprintf("%x", h.Sum(nil))
}
