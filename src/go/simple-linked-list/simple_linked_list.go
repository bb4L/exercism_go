package linkedlist

import (
	"errors"
)

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
	for elem := l.firstElement; elem != nil; elem = elem.next {
		size++
	}
	return size
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
	for elem := l.firstElement; elem != nil; elem = elem.next {
		result = append([]int{elem.data}, result...)
	}
	return result
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
