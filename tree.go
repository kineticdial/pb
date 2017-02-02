package pb

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

type Tree struct {
	Refs []*TreeRef
}

func GetTree(k string) (*Tree, error) {
	contents, err := GetObject(k)
	if err != nil {
		return nil, err
	}

	// Trim potential new line from end of file and split lines.
	refs := strings.Split(
		strings.TrimRight(contents, "\n"),
		"\n",
	)

	trs := make([]*TreeRef, 0, len(refs))
	for _, ref := range refs {
		decodedRef, err := DecodeTreeRef(ref)
		if err != nil {
			return nil, err
		}

		trs = append(trs, decodedRef)
	}

	t := &Tree{trs}
	return t, nil
}

func (t *Tree) Put() error {
	return PutObject(t)
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
