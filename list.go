package main

import "errors"

type list struct {
	next *list
	data int
}

func NewList(size int, deepness int) (*list, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	if size == 1 {
		return &list{data: deepness}, nil
	}
	nextList, err := NewList(size-1, deepness+1)
	if err != nil {
		return nil, err
	}
	return &list{data: deepness, next: nextList}, nil
}

func (l *list) Add(data int) {
	if l.next == nil {
		l.next = &list{data: data}
		return
	}
	l.next.Add(data)
}

func (l *list) Get(index int) int {
	if index == 0 {
		return l.data
	}
	return l.next.Get(index - 1)
}

// deepness is the number of nodes to go back
// if deepness is 0, the current node will be removed
// if deepness is 1, the previous node will be removed
// if deepness is 2, the node before the previous will be removed
// and so on
func (l *list) Remove(deepness int, previous ...*list) {
	// if previous is not empty, it means that the current node is not the first
	// so we need to remove the current node and link the previous node to the next node
	var prev *list
	if len(previous) > 0 {
		prev = previous[0]
	}
	if deepness == 0 {
		if prev != nil {
			prev.next = l.next
			return
		} else {
			*l = *l.next
			return
		}
	}
	l.next.Remove(deepness-1, l)
}

func (l *list) Length() int {
	if l.next == nil {
		return 1
	}
	return 1 + l.next.Length()
}

func (l *list) PrintAll() {
	println(l.data)
	if l.next == nil {
		return
	}
	l.next.PrintAll()
}
