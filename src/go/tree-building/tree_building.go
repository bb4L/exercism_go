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

	for k, i := range data {
		if i.ID != k || (i.ID < i.Parent || (i.ID == i.Parent && i.ID != 0)) || (indices[i.ID] != 0) {
			return nil, errors.New("Incorrec structure")
		}

		m[i.ID] = Node{ID: i.ID}
		indices[i.ID] = k

	}

	for index := range m {
		n := &m[index]

		parent := data[indices[index]].Parent

		if parent != index {
			m[parent].Children = append(m[parent].Children, n)
		}

	}

	return &m[0], nil
}
