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

	m := make([]Node, len(data))
	indices := make([]int, len(data))
	rootNodeSet := false

	for k, i := range data {
		if i.ID >= len(data) || i.ID < 0 {
			return nil, errors.New("Incorrect ID")
		}

		if i.ID < i.Parent || (i.ID == i.Parent && i.ID != 0) {
			return nil, errors.New("Node id has to be strictly lower than the parent")
		}

		if indices[i.ID] != 0 {
			return nil, errors.New("Node set twice ")
		}

		m[i.ID] = Node{ID: i.ID}
		indices[i.ID] = k

		if i.ID == 0 {
			if rootNodeSet {
				return nil, errors.New("Root node set twice")
			}
			rootNodeSet = true
		}
	}

	if !rootNodeSet {
		return nil, errors.New("No root node")
	}

	for index := range m {
		n := &m[index]

		parent := data[indices[index]].Parent

		if parent != index {
			m[parent].Children = append(m[parent].Children, n)
			sort.Slice(m[parent].Children, func(i, j int) bool {
				return m[parent].Children[i].ID < m[parent].Children[j].ID
			})
		}

	}

	return &m[0], nil
}
