// Package main offers an implementation of the quicksort sorting algorithm.
package main

import (
	"fmt"
	r "math/rand"
	"time"
)

// Text assignment
// Implement the quicksort algorithm in Go
// and sort some randomly-generated numeric data.

const capacity int = 100

func partition(A *[capacity]int, p, r int) int {
	x := A[r]
	for j := p; j < r; j++ {
		if A[j] < x {
			A[j], A[p] = A[p], A[j]
			p++
		}
	}
	A[p], A[r] = A[r], A[p]
	return p
}

func quicksort(A *[capacity]int, p, r int) {
	if p > r {
		return
	}
	q := partition(A, p, r)
	quicksort(A, p, q-1)
	quicksort(A, q+1, r)
}

// QuickSort sorts an input given array.
func QuickSort(A *[capacity]int) {
	quicksort(A, 0, len(A)-1)
}

func main() {

	// Initializing the Seed for random numbers
	r.Seed(time.Now().UnixNano())

	A := new([capacity]int)

	for i := 0; i < capacity; i++ {
		A[i] = r.Intn(capacity)
	}

	QuickSort(A)
	fmt.Println("This is the array after sorting procedure:", A)
}
