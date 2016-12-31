package obj

import (
	"crypto/sha1"
	"fmt"

	"github.com/Lead-SCM/pb/db"
)

type Tree struct {
	Trees map[string]string
	Blobs map[string]string
}

func NewTree(Trees map[string]string, Blobs map[string]string) *Tree {
	return &Tree{Trees, Blobs}
}

func GetTree(k string) (*Tree, error) {
	var t Tree
	err := db.Get(k, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Tree) Put() error {
	err := db.Put(t.Hash(), t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tree) Hash() string {
	h := sha1.New()
	for name, hash := range t.Trees {
		h.Write([]byte(name))
		h.Write([]byte(hash))
	}

	for name, hash := range t.Blobs {
		h.Write([]byte(name))
		h.Write([]byte(hash))
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
