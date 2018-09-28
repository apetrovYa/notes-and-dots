package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args

	if len(arguments) < 1 {
		log.Println("Provide any argument")
		os.Exit(1)
	}

	filename := arguments[1]
	fileinfo, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}
	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, err := filepath.EvalSymlinks(filename)
		if err == nil {
			fmt.Println(realpath)
		}
	}
}
