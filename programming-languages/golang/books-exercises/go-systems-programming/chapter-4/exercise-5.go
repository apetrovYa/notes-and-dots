package main

import (
	"fmt"
	"log"
)

// Test assignment
// The implementation of tree.go is far from complete! Try to implement a
// function that checks whether a value can be found in the tree and another
// function that allows you to delete a tree node.

type node struct {
	value int
	left  *node
	right *node
}

func newTreeNode(v int) *node {
	return &node{v, nil, nil}
}

func countNodes(t *node) int {
	if t == nil {
		return 0
	}
	return 1 + countNodes(t.left) + countNodes(t.right)
}

func insert(t *node, i *node) *node {
	if i == nil {
		log.Println("The node to insert is nil")
		return t
	}

	if t == nil {
		log.Println("The tree is empty")
		return i
	}

	if t.left == nil && t.right == nil {
		t.left = i
		return t
	}

	if t.left == nil && t.right != nil {
		t.left = i
		return t
	}
	if t.left != nil && t.right == nil {
		t.right = i
		return t
	}

	if countNodes(t.left) < countNodes(t.right) {
		t.left = insert(t.left, i)
	} else {
		t.right = insert(t.right, i)
	}
	return t
}

func printPreOrder(t *node) {
	if t != nil {
		fmt.Println("Value:", t.value)
		printPreOrder(t.left)
		printPreOrder(t.right)

	}
}

func printInOrder(t *node) {
	if t != nil {
		printPreOrder(t.left)
		fmt.Println(t.value)
		printPreOrder(t.right)
	}
}
func printPostOrder(t *node) {
	if t != nil {
		printPreOrder(t.left)
		printPreOrder(t.right)
		fmt.Println(t.value)
	}
}

func searchValue(t *node, v int) *node {
	if t == nil {
		log.Print("The tree is empty")
		return nil
	}
	if t.value == v {
		log.Print("Item found. Returning the pointer to the node containing it.")
		return t
	}
	x := searchValue(t.left, v)
	if x != nil {
		return x
	}
	return searchValue(t.right, v)
}

func main() {

	var tree *node

	for i := 0; i < 10; i++ {
		tree = insert(tree, newTreeNode(i))
	}

	printPreOrder(tree)

	y := searchValue(tree, -1)

	if y != nil {
		fmt.Println("The founded item has the following address and value", &y, y.value)
	}

}
