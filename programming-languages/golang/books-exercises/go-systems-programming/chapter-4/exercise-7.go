package main

import (
	"fmt"
	"log"
)

// Test assignment
// Once again, the hash table implementation of hash.go is incomplete as it
// allows duplicate entries. So, implement a function that searches the hash
// table for a key before inserting it.

type node struct {
	value int
	next  *node
}

type hashTable struct {
	size  int
	table map[int]*node
}

func hash(v, size int) int {
	return v % size
}

func searchValue(l *node, v int) bool {
	if l == nil {
		return false
	}

	if l.value == v {
		return true
	}
	return searchValue(l.next, v)
}

func insert(h *hashTable, v int) int {
	index := hash(v, h.size)
	element := &node{value: v, next: h.table[index]}
	h.table[index] = element
	return index
}

func insert2(h *hashTable, v int) int {
	index := hash(v, h.size)
	found := searchValue(h.table[index], v)
	if !found {
		return insert(h, v)
	}
	log.Print("Item already present inside the HashTable")
	return -1
}

func main() {

	const s int = 5
	h := &hashTable{s, make(map[int]*node)}

	for i := 0; i < s; i++ {
		j := insert2(h, i)
		fmt.Println(j)
	}

	j := insert2(h, 2)
	fmt.Println(j)

}
