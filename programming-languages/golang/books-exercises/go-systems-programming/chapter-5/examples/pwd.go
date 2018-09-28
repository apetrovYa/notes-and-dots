//PWD(1)
//
//NAME
//       pwd - print name of current/working directory
//
//SYNOPSIS
//       pwd [OPTION]...
//
//DESCRIPTION
//       Print the full filename of the current working directory.
//
//       -L, --logical
//              use PWD from environment, even if it contains symlinks
//
//       -P, --physical
//              avoid all symlinks
//
//       --help display this help and exit
//
//       --version
//              output version information and exit
//
//       If no option is specified, -P is assumed.
//
//       NOTE: your shell may have its own version of pwd, which usually
//       supersedes the version described here.  Please refer to your shell's
//       documentation for details about the options it supports.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const v string = "0.2"

func usage() {
	fmt.Println("NAME")
	fmt.Println("pwd - print name of current/working directory")
	fmt.Println()
	fmt.Println("SYNOPSIS")
	fmt.Println("\t pwd [OPTION]...")
	fmt.Println()
	fmt.Println("DESCRIPTION")
	fmt.Println("\t  Print the full filename of the current working directory.")
	fmt.Println("\t -L, --logical")
	fmt.Println("\t \t use PWD from environment, even if it contains symlinks")
	fmt.Println("\t -P, --physical")
	fmt.Println("\t \t  avoid all symlinks.")
	fmt.Println("\t --help display this help and exit")
	fmt.Println("\t --version")
	fmt.Println("\t \t  output version information and exit")
	fmt.Println("If no option is specified, -L is assumed.")
}

func logical(fpath string) {
	fmt.Println(fpath)
}

func physical(fpath string) {
	fileinfo, _ := os.Lstat(fpath)
	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, _ := filepath.EvalSymlinks(fpath)
		fmt.Println(realpath)
	} else {
		fmt.Println(fpath)
	}
}

func version() {
	fmt.Printf("pwd (GO utils) %s\n", v)
	fmt.Println("This is free software: you are free to change and redistribute it.")
	fmt.Println("There is NO WARRANTY, to the extent permitted by law.")
	fmt.Println("Written by Andrei Petrov.")
}

func notValidOption() {
	fmt.Println("Inserted option is not valid.")
	fmt.Println()
}

func manageOption(option string) {
	filename, err := os.Getwd()

	if err != nil {
		log.Fatal("Something bad happened while getting the current working directory.")
	}

	switch option {
	case "-L", "--logical":
		logical(filename)
	case "-P", "--physical":
		physical(filename)
	case "-h", "--help":
		usage()
	case "--version":
		version()
	default:
		notValidOption()
	}

}

func main() {

	switch len(os.Args) {
	case 2:
		manageOption(os.Args[1])
	default:
		manageOption("-L")
	}
}
