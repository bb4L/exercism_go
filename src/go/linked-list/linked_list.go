package linkedlist

import (
	"errors"
)

// ErrEmptyList empty list error
var ErrEmptyList = errors.New("empty list")

// Entry one entry of the list
type Entry struct {
	Val       interface{}
	nextEntry *Entry
	prev      *Entry
}

// Next pointer to the next node
func (entry Entry) Next() *Entry {
	return entry.nextEntry
}

// Prev pointer to the previous node
func (entry Entry) Prev() *Entry {
	return entry.prev
}

// List struct
type List struct {
	firstEntry *Entry
	lastEntry  *Entry
}

// NewList creates a new linked list preserving the order of the values
func NewList(values ...interface{}) *List {
	list := List{}

	for _, v := range values {
		list.PushBack(v)
	}

	return &list
}

// PushFront insert value at the front of the list
func (list *List) PushFront(value interface{}) {
	previousFirst := list.firstEntry
	list.firstEntry = &Entry{Val: value, nextEntry: previousFirst}
	if previousFirst == nil {
		list.lastEntry = list.firstEntry
		return
	}
	previousFirst.prev = list.firstEntry
}

// PushBack insert value at the back of the list
func (list *List) PushBack(value interface{}) {
	var oldLast *Entry

	if list.lastEntry != nil {
		oldLast = list.lastEntry
	}

	list.lastEntry = &Entry{Val: value, prev: oldLast}

	if oldLast != nil {
		oldLast.nextEntry = list.lastEntry
	} else {
		list.firstEntry = list.lastEntry
	}
}

// PopFront remove value from the front of the list
func (list *List) PopFront() (interface{}, error) {
	if list.firstEntry == nil {
		return nil, ErrEmptyList
	}

	returnValue := list.firstEntry.Val
	list.firstEntry = list.firstEntry.nextEntry

	if list.firstEntry != nil {
		list.firstEntry.prev = nil
	} else {
		list.lastEntry = nil
	}

	return returnValue, nil
}

// PopBack remove value from the back of the list
func (list *List) PopBack() (interface{}, error) {
	if list.lastEntry == nil {
		return nil, ErrEmptyList
	}

	returnValue := list.lastEntry.Val
	list.lastEntry = list.lastEntry.prev

	if list.lastEntry != nil {
		list.lastEntry.nextEntry = nil
	} else {
		list.firstEntry = nil
	}

	return returnValue, nil
}

// First returns a pointer to the first node
func (list *List) First() *Entry {
	return list.firstEntry
}

// Last returns a pointer to the last node
func (list List) Last() *Entry {
	return list.lastEntry
}

// Reverse reverse the linked list
func (list *List) Reverse() {
	entry := list.firstEntry
	list.firstEntry = list.lastEntry
	list.lastEntry = entry

	if entry == nil {
		return
	}
	nextEntry := entry.nextEntry

	for {
		entry.prev, entry.nextEntry = entry.nextEntry, entry.prev

		if entry == list.firstEntry {
			break
		}
		entry, nextEntry = nextEntry, nextEntry.nextEntry
	}
}
