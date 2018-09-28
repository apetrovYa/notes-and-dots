//WHICH(1)
//
//NAME
//       which - locate a command
//
//SYNOPSIS
//       which [-a] filename ...
//
//DESCRIPTION
//       which  returns  the  pathnames of the files (or links) which would be executed in the current environment,
//       had its arguments been given as commands in a strictly POSIX-conformant shell.  It does this by  searching
//       the PATH for executable files matching the names of the arguments. It does not canonicalize path names.
//
//OPTIONS
//       -a     print all matching pathnames of each argument
//
//EXIT STATUS
//       0      if all specified commands are found and executable
//
//       1      if one or more specified commands is nonexistent or not executable
//
//       2      if an invalid option is specified
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func searchBinaryFile(path string, executable string) string {
	files, _ := ioutil.ReadDir(path)

	for _, f := range files {
		filename := f.Name()
		if filename == executable {
			return filename
		}
	}

	return ""
}

func insertBinaryItem(list []string, e string) []string {
	b := false

	for _, v := range list {
		if v == e {
			b = true
		}
	}

	if !b {
		list = append(list, e)
	}
	return list

}

func searchExecutable(executables []string, executable string) string {
	for _, path := range executables {
		b := searchBinaryFile(path, executable)

		if b != "" {
			return path + "/" + b

		}
	}
	return ""
}

func expandEnvironmentVariable(e string) string {
	return os.ExpandEnv(e)
}

func getExecutableList(envPath string) []string {
	if envPath != "" {
		return strings.Split(envPath, ":")
	}
	log.Printf("The PATH variable is empty")
	return nil
}

func isExecutable(file string) bool {
	f, _ := os.Stat(file)
	if f.Mode()&0111 != 0 {
		return true
	}
	return false
}

func main() {
	args := os.Args

	if len(args) == 1 {
		os.Exit(0)
	}

	var commands []string

	if args[1] == "-a" {
		commands = append(commands, args[2:]...)
	} else {
		commands = append(commands, args[1])
	}

	output := map[string]bool{}

	e := expandEnvironmentVariable("$PATH")
	e2 := getExecutableList(e)

	for _, s := range commands {
		e3 := searchExecutable(e2, s)
		output[e3] = isExecutable(e3)
	}

	for _, v := range output {
		if v == false {
			os.Exit(1)
		}
	}

	for s := range output {
		fmt.Println(s)
	}
	os.Exit(0)
}
