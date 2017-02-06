package lib

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

// Tree represents a collection of TreeRefs as an abstraction of a directory and
// its contents.
type Tree struct {
	Refs []*TreeRef
}

// GetTree fetches a Tree from the file k/v store by its SHA1 hash.
func GetTree(k string) (*Tree, error) {
	contents, err := getObject(k)
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

// Put stores a Tree in memory into the file k/v store.
func (t *Tree) Put() error {
	return putObject(t)
}

// String returns a Tree's contents.
func (t *Tree) String() string {
	buf := bytes.NewBufferString("")

	sort.Sort(ByString(t.Refs))

	for _, tref := range t.Refs {
		buf.WriteString(tref.String())
		buf.WriteString("\n")
	}

	return buf.String()
}

// Hash calculates a SHA1 hash of the Tree's contents.
func (t *Tree) Hash() string {
	h := sha1.New()
	h.Write([]byte(t.String()))

	return fmt.Sprintf("%x", h.Sum(nil))
}
