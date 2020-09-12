package tree

import (
	"fmt"
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

	m := map[int]*Node{}

	for i, r := range data {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, fmt.Errorf("bad record: %v (out of sequence or bad parent ID)", r)
		}

		m[r.ID] = &Node{ID: r.ID}

		parent := r.Parent
		if r.ID > 0 {
			m[parent].Children = append(m[parent].Children, m[i])
		}
	}

	return m[0], nil
}
