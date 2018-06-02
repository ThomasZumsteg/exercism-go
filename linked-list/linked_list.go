package linkedlist

import "errors"

type List struct {
	head *Element
	tail *Element
}

type Element struct {
	Val  interface{}
	next *Element
	prev *Element
}

var ErrEmptyList = errors.New("Empty List")

func NewList(items ...interface{}) *List {
	ll := List{}
	for _, item := range items {
		ll.PushBack(item)
	}
	return &ll
}

func (list *List) First() *Element {
	return list.head
}

func (list *List) Last() *Element {
	return list.tail
}

func (element *Element) Prev() *Element {
	return element.prev
}

func (element *Element) Next() *Element {
	return element.next
}

func (list *List) PushBack(item interface{}) {
	e := Element{item, nil, list.tail}
	if list.tail == nil {
		list.head = &e
	} else {
		list.tail.next = &e
	}
	list.tail = &e
}

func (list *List) PushFront(item interface{}) {
	e := Element{item, list.head, nil}
	if list.head == nil {
		list.tail = &e
	} else {
		list.head.prev = &e
	}
	list.head = &e
}

func (list *List) PopBack() (interface{}, error) {
	element := list.Last()
	if element == nil {
		return 0, ErrEmptyList
	}
	list.tail = element.Prev()
	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.next = nil
	}
	return element.Val, nil
}

func (list *List) PopFront() (interface{}, error) {
	element := list.First()
	if element == nil {
		return 0, ErrEmptyList
	}
	list.head = element.Next()
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
	return element.Val, nil
}

func (list *List) Reverse() {
	for e := list.First(); e != nil; e = e.Prev() {
		e.next, e.prev = e.prev, e.next
	}
	list.head, list.tail = list.tail, list.head
}
