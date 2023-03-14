// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Path Searcher
//
//  Your program should search inside the path environment
//  variable.
//
//  Remove the corpus constant then get the corpus from the
//  environment variable "Path" or "PATH".
//
// HINTS
//  1. Search the web to find out what is an environment
//     variable and how to use it (if you don't know
//     what it is already).
//
//  2. Look up for the necessary functions for getting
//     an environment variable. It's in the "os" package.
//
//     Search for it on the Go online documentation.
//
//  3. Look up for the necessary function for splitting
//     the path variable into directory names.
//
//     It's in the "path/filepath" package.
//
// EXAMPLE
//  For example, on my Mac, my PATH environment variable
//  looks like this:
//
//    "/usr/local/bin:/sbin:/Users/inanc/go/bin"
//
//  So, if the user runs the program like this:
//
//    go run main.go /sbin
//
//  It should print this:
//
//    #2 : "/sbin"
// ---------------------------------------------------------

// ---------------------------------------------------------
// BONUS EXERCISE
//  Make your program cross platform. So, it can search
//  the path environment variable when you run it on
//  a Windows or on a Mac (OS X) or on a Linux.
//
// BONUS EXERCISE #2
//  Also find the directories for the directories that
//  contains the searched query.
//
//  And let it match in a case-insensitive manner. See the examples.
//
//  EXAMPLE:
//    Let's say:
//      PATH="/usr/local/bin:/sbin:/Users/inanc/go/bin"
//
//  So, if the user runs the program like this:
//    go run main.go bin
//
//  It should print this:
//    #1 : "/usr/local/bin"
//    #2 : "/sbin"
//    #3 : "/Users/inanc/go/bin"
//
//  Or like this (case insensitive):
//    go run main.go Us
//
//  Output:
//    #1 : "/usr/local/bin"
//    #2 : "/Users/inanc/go/bin"
// ---------------------------------------------------------

func main() {
	// get the arguments from the commandline
	args := os.Args[1:]
	// get the path to the PATH variable
	p := os.Getenv("PATH")
	// split the path
	words := filepath.SplitList(p)

	// loop over the args
	for _, a := range args {
		// loop over the paths
		for i, w := range words {
			// lower case the args and the paths
			a, w = strings.ToLower(a), strings.ToLower(w)
			// check if the path contains the arg
			if strings.Contains(w, a) {
				fmt.Printf("#%-2d : %s\n", i+1, w)
			}
		}
	}
}
