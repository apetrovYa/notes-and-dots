package main

import (
	"fmt"
	"log"
)

// Test assignment
// Implement a doubly linked list.

type node struct {
	value int
	next  *node
	prev  *node
}

func newNode(v int) *node {
	return &node{v, nil, nil}
}

func insert(l *node, n *node) *node {
	if n == nil {
		log.Println("The new node is nil")
		return l
	}
	if l == nil {
		log.Println("The main list is empty")
		return n
	} else {
		log.Println("Going to insert the new node")
		n.next = l
		l.prev = n
		return n
	}
}

func printList(l *node) {
	fmt.Println("Going to print the list in forward direction")
	for l != nil {
		fmt.Print(l.value, "-> ")
		l = l.next
	}
	fmt.Println()
}

func printListInReverseOrder(l *node) {
	fmt.Println("Going to print the list in reverse direction")
	if l == nil {
		log.Println("The list to print in reverse order is empty")
		return
	}

	for l.next != nil {
		l = l.next
	}

	for l != nil {
		fmt.Print(l.value, " <-")
		l = l.prev
	}
	fmt.Println()
}

func main() {

	var list *node
	//list = nil

	for i := 0; i < 10; i++ {
		list = insert(list, newNode(i))
	}
	printList(list)
	printListInReverseOrder(list)

}
