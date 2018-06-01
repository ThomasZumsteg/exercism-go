package linkedlist

import "errors"

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(items []int) *List {
	var head = List{nil, 0}
	for _, item := range items {
		head.Push(item)
	}
	return &head
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Push(data int) {
	list.head = &Element{data, list.head}
	list.size++
}

func (list *List) Pop() (data int, err error) {
	if list.head == nil {
		err = errors.New("Empty list")
	} else {
		data = list.head.data
		list.head = list.head.next
	}
	list.size--
	return
}

func (list *List) Array() []int {
	result := make([]int, list.Size())
	for i, head := list.Size()-1, list.head; head != nil; i, head = i-1, head.next {
		result[i] = head.data
	}
	return result
}

func (list *List) Reverse() *List {
	result := New([]int{})
	for item, err := list.Pop(); err == nil; item, err = list.Pop() {
		result.Push(item)
	}
	return result
}
