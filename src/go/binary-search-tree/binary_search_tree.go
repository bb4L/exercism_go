// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package binarysearchtree should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package binarysearchtree

// SearchTreeData is a node of the searchtree
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// NewBst creates and returns a new SearchTreeData.
func NewBst(i int) SearchTreeData {
	return SearchTreeData{data: i}
}

// Insert inserts an int into the SearchTreeData.
// Inserts happen based on the rules of a BinarySearchTree
func (std *SearchTreeData) Insert(i int) {
	nodeToCheck := std

	for {
		if nodeToCheck.data >= i {
			if nodeToCheck.left == nil {
				nodeToCheck.left = &SearchTreeData{data: i}
				return
			}
			nodeToCheck = nodeToCheck.left
		} else {
			if nodeToCheck.right == nil {
				nodeToCheck.right = &SearchTreeData{data: i}
				return
			}
			nodeToCheck = nodeToCheck.right
		}
	}
}

// MapString returns the ordered contents of SearchTreeData as a []string.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []string ["1", "3", "5", "7"].
func (std *SearchTreeData) MapString(fun func(int) string) (result []string) {
	if std.left != nil {
		result = append(std.left.MapString(fun), result...)
	}
	result = append(result, fun(std.data))
	if std.right != nil {
		result = append(result, std.right.MapString(fun)...)
	}
	return
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(mapping func(int) int) (result []int) {
	if std.left != nil {
		result = append(std.left.MapInt(mapping), result...)
	}
	result = append(result, mapping(std.data))
	if std.right != nil {
		result = append(result, std.right.MapInt(mapping)...)
	}
	return
}
