package tree

import (
	"errors"
	"sort"
)

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

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	m := make([]Node, len(data))
	indices := make([]int, len(data))

	for i, r := range data {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("Incorrect structure")
		}

		m[r.ID] = Node{ID: r.ID}
		indices[r.ID] = i

		n := &m[i]

		parent := data[indices[i]].Parent

		if parent != i {
			m[parent].Children = append(m[parent].Children, n)
		}
	}

	return &m[0], nil
}
