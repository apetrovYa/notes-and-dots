package main

import (
	"fmt"
	"sort"
)

// Text assignment:
// Create your own structure, make a slice and use the sort.Slice() to sort
// the elements of the slice you created.

type book struct {
	name       string
	pages      int
	references int
}

func main() {

	books := make([]book, 0)
	books = append(books, book{"Go programming", 490, 10})
	books = append(books, book{"Go Systems programming", 900, 1})
	books = append(books, book{"Mysql Administration", 1000, 99})

	fmt.Println("Printing the slice content\n", books)

	sort.Slice(books, func(i, j int) bool {
		return books[i].pages < books[j].pages
	})
	fmt.Println("Printing the sorted slice (<)", books)

	sort.Slice(books, func(i, j int) bool {
		return books[i].pages > books[j].pages
	})
	fmt.Println("Printing the sorted slice (>)", books)
}
