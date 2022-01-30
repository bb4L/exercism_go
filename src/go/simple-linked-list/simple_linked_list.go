package linkedlist

import (
	"errors"
)

// List own implementation of a list
type List struct {
	firstElement *Element
}

// A Element of a list
type Element struct {
	data int
	next *Element
}

// New creates a new list from a slice
func New(data []int) *List {
	list := &List{}
	for _, val := range data {
		list.Push(val)
	}
	return list
}

// Size returns the list size
func (l *List) Size() (size int) {
	for elem := l.firstElement; elem != nil; elem = elem.next {
		size++
	}
	return size
}

// Push a element to the list
func (l *List) Push(element int) {
	l.firstElement = &Element{data: element, next: l.firstElement}
}

// Pop a element from the list
func (l *List) Pop() (int, error) {
	returnElement := l.firstElement
	if returnElement == nil {
		return -1, errors.New("cannot pop on empty")
	}
	l.firstElement = returnElement.next

	return returnElement.data, nil
}

// Array converts the list to a array
func (l *List) Array() (result []int) {
	for elem := l.firstElement; elem != nil; elem = elem.next {
		result = append([]int{elem.data}, result...)
	}
	return result
}

// Reverse the list
func (l *List) Reverse() *List {
	if l.firstElement == nil {
		return l
	}

	oldelem := l.firstElement
	newElem := &Element{data: oldelem.data}

	for oldelem.next != nil {
		oldelem = oldelem.next
		newElem = &Element{data: oldelem.data, next: newElem}
	}
	return &List{firstElement: newElem}
}
