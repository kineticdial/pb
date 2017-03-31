package lib

import "fmt"
import "strconv"
import "strings"

// TreeRef represents the n nodes (blobs or additional trees) that a Tree can
// reference.
type TreeRef struct {
	Perms   int    // Permissions of referenced file/dir
	RefType string // Blob or Tree
	Name    string // Name of file/dir
	Ref     object // Blob/Tree
}

// DecodeTreeRef takes a raw tab-delimited string and serializes it into a
// TreeRef struct.
func DecodeTreeRef(s string) (*TreeRef, error) {
	items := strings.Split(s, "\t")

	// Parse int as octal
	perms, err := strconv.ParseInt(items[0], 8, 64)
	if err != nil {
		return nil, err
	}

	refType := items[1]
	var ref object

	if refType == "Tree" {
		ref, err = GetTree(items[3])
	} else {
		ref, err = GetBlob(items[3])
	}

	if err != nil {
		return nil, err
	}

	return &TreeRef{
		Perms:   int(perms),
		RefType: items[1],
		Name:    items[2],
		Ref:     ref,
	}, nil
}

// String returns a tab-delimited representation of a TreeRef. File
// permissions are represented as octals.
func (tr *TreeRef) String() string {
	return fmt.Sprintf(
		"%o\t%s\t%s\t%s",
		tr.Perms,
		tr.RefType,
		tr.Name,
		tr.Ref.Hash(),
	)
}

// ByString sorts TreeRefs by their string representation.
type ByString []*TreeRef

func (a ByString) Len() int           { return len(a) }
func (a ByString) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByString) Less(i, j int) bool { return a[i].String() > a[j].String() }
