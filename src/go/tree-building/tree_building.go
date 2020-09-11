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

	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	m := make([]Node, len(data))

	for i, r := range data {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("Incorrect structure")
		}

		m[r.ID] = Node{ID: r.ID}

		parent := r.Parent
		if parent != i {
			m[parent].Children = append(m[parent].Children, &m[i])
		}
	}

	return &m[0], nil
}
