package algorithms

import "github.com/lead-scm/pb/lib"

// BuildTreeFromWorkingIndex processes the working index into an in-memory Tree
// and then persists it onto the filesystem. It returns the root Tree.
// Calling code is responsible for reading the working index file and splitting
// each path by "/".
func BuildTreeFromWorkingIndex(index [][]string) (*lib.Tree, error) {
	root := &lib.Tree{}

	for _, fp := range index {
		err := buildTreeFromPath(root, fp)
		if err != nil {
			return nil, err
		}
	}

	root.Put()
	return root, nil
}

// Given a path to a file (and its hash) we'll need to form a tree and merge
// it onto the tree passed in as "root".
//
// root *lib.Tree -- the root of the tree to merge onto
// filepath []string -- a representation of a filepath to a blob, e.g.:
//       Dir     Dir       File            Hash
//     ["test", "models", "user_test.go", "0blc20393bola90a3932b309b209e0thh"]
func buildTreeFromPath(root *lib.Tree, filepath []string) error {
	path := filepath[:len(filepath)-2]
	filename := filepath[len(filepath)-2 : len(filepath)-1][0]
	hash := filepath[len(filepath)-1:][0]
	currentNode := root

	for _, node := range path {
		found := false
		for _, tr := range currentNode.Refs {
			if node == tr.Name {
				found = true
				currentNode, _ = tr.Ref.(*lib.Tree)
			}
		}

		if !found {
			tree := &lib.Tree{}
			tr := &lib.TreeRef{
				Perms:   0700, // TODO
				RefType: "Tree",
				Name:    node,
				Ref:     tree,
			}
			currentNode.Refs = append(currentNode.Refs, tr)
			currentNode = tree
		}
	}

	blob, err := lib.GetBlob(hash)
	if err != nil {
		return err
	}
	tr := &lib.TreeRef{
		Perms:   0700, // TODO
		RefType: "Blob",
		Name:    filename,
		Ref:     blob,
	}
	currentNode.Refs = append(currentNode.Refs, tr)

	return nil
}
