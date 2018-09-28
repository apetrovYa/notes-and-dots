package main

import (
	"flag"
	"fmt"
)

func main() {
	minusC := flag.Bool("c", false, "First argument")
	minusA := flag.Bool("a", true, "Second argument")
	minusK := flag.Int("k", 0, "Third argument")

	flag.Parse()

	fmt.Println(*minusC)
	fmt.Println(*minusA)
	fmt.Println(*minusK)

	for i, f := range flag.Args() {
		fmt.Println(i, ":", f)
	}
}
