package tree

import "errors"

// Record represents a record to setup the tree
type Record struct {
	ID     int
	Parent int
}

// Node a node for the tree
type Node struct {
	ID       int
	Children []*Node
}

// Build a tree from the given records
func Build(data []Record) (*Node, error) {
	if len(data) == 0 {
		return nil, nil
	}

	m := make(map[int]Node)

	for _, i := range data {
		m[i.ID] = Node{ID: i.ID}
	}

	rootNode := Node{}
	rootNodeSet := false

	for _, i := range data {
		n, ok := m[i.ID]

		if !ok {
			return nil, errors.New("Could not find node")
		}

		parent, ok := m[i.Parent]

		if !ok {
			return nil, errors.New("Could not find parent")
		}

		if i.Parent != i.ID {
			parent.Children = append(parent.Children, &n)
		} else {
			if rootNodeSet {
				return nil, errors.New("Can not have two root nodes")
			}
			rootNode = n
			rootNodeSet = true
		}

	}

	return &rootNode, nil
}
