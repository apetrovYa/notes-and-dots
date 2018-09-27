package main

import (
	"fmt"
)

// Test assignment
// Similarly, the implementation of the linkedList.go file is also incomplete.
// Try to implement a function for deleting a node and another one for inserting
// a node somewhere inside the linked list.

type node struct {
	value int
	next  *node
}

func insert(l *node, v int) *node {
	if l == nil {
		return &node{v, nil}
	}

	if l.next == nil {
		l.next = &node{v, nil}
		return l
	}

	l.next = insert(l.next, v)
	return l
}

func searchValue(l *node, v int) bool {
	if l != nil {
		if l.value == v {
			return true
		}
		return searchValue(l.next, v)
	}
	return false
}

func deleteItem(l *node, v int) *node {
	if l != nil {
		if l.value == v {
			t := l.next
			l.next = nil
			return t
		}
		l.next = deleteItem(l.next, v)
		return l
	}
	return nil
}

func traverse(list *node) {

	for list != nil {
		fmt.Print(list.value, " ")
		list = list.next
	}
	fmt.Println()
}

func main() {
	var li *node
	for i := 0; i < 10; i++ {
		li = insert(li, i)
	}

	l := deleteItem(li, 3)
	traverse(l)
}
