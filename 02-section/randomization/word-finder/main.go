package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" +
	"lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	args := os.Args[1:]

queries:
	for _, q := range args {
		for i, w := range words {
			if strings.EqualFold(q, w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break queries
				// continue queries
			}
		}
	}
}
