package algorithms

import "github.com/lead-scm/pb/lib"

func BuildTreeForCommit(index [][]string) (*lib.Tree, error) {
	root := &lib.Tree{}

	for _, fp := range index {
		err := toTree(root, fp)
		if err != nil {
			return nil, err
		}
	}

	root.Put()
	return root, nil
}

func toTree(root *lib.Tree, filepath []string) error {
	path := filepath[:len(filepath)-2]
	filename := filepath[len(filepath)-2 : len(filepath)-1]
	hash := filepath[len(filepath)-1:]
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
				Name:    node, // Will always be slice of size 1
				Ref:     tree,
			}
			currentNode.Refs = append(currentNode.Refs, tr)
			currentNode = tree
		}
	}

	blob, err := lib.GetBlob(hash[0])
	if err != nil {
		return err
	}
	tr := &lib.TreeRef{
		Perms:   0700, // TODO
		RefType: "Blob",
		Name:    filename[0], // Will always be slice of size 1
		Ref:     blob,
	}
	currentNode.Refs = append(currentNode.Refs, tr)

	return nil
}
