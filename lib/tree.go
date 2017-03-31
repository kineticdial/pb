package lib

import "bytes"
import "crypto/sha256"
import "fmt"
import "sort"
import "strings"

// Tree represents a collection of TreeRefs as an abstraction of a directory and
// its contents.
type Tree struct {
	Refs []*TreeRef // Files/subdirs of current dir
}

// GetTree fetches a Tree from the file k/v store by its SHA256 hash.
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

// Put stores a in memory Tree (and all of its children) into the file
// k/v store.
func (t *Tree) Put() error {
	for _, tr := range t.Refs {
		tr.Ref.Put()
	}
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

// Hash calculates a SHA256 hash of the Tree's contents.
func (t *Tree) Hash() string {
	h := sha256.New()
	h.Write([]byte(t.String()))

	return fmt.Sprintf("%x", h.Sum(nil))
}
