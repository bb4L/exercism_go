package linkedlist

import (
	"errors"
)

// Define the List and Element types here.
type List struct {
	firstElement *Element
}

type Element struct {
	data int
	next *Element
}

func New(data []int) *List {
	list := &List{}
	for _, val := range data {
		list.Push(val)
	}
	return list
}

func (l *List) Size() (size int) {
	elem := l.firstElement
	for elem != nil {
		size++
		elem = elem.next
	}
	return
}

func (l *List) Push(element int) {
	l.firstElement = &Element{data: element, next: l.firstElement}
}

func (l *List) Pop() (int, error) {
	returnElement := l.firstElement
	if returnElement == nil {
		return -1, errors.New("cannot pop on empty")
	}
	l.firstElement = returnElement.next

	return returnElement.data, nil
}

func (l *List) Array() (result []int) {
	elem := l.firstElement
	for elem != nil {
		result = append([]int{elem.data}, result...)
		elem = elem.next
	}
	return
}

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
