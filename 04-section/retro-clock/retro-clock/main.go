// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Refactor
//
//	Goal is refactoring the clock project by moving some of its parts to
//	a new file in the same package.
//
//	1. Create a new file: placeholders.go
//	2. Move the placeholder type to placeholder.go
//	3. Move all the placeholders (zero to nine and the colon) to placeholder.go
//	4. Move the digits array to placeholders.go
//
// HINT
//
//   - placeholders.go file should belong to main package as well
//
//     To remember how to do so: Rewatch the "PART I — Packages, Scopes and Importing"
//     section.
//
//   - Short declaration won't work in the package scope.
//     Remember why by watching: "Short Declaration: Package Scope" lecture
//
//   - If you receive the following error and you don't know what to do watch:
//     "Packages - Learn how to run multiple files" lecture.
//
//     # command-line-arguments
//     undefined: placeholder
//     undefined: colon
//
// EXPECTED OUTPUT
//
//	The same output before the refactoring.
//
// ---------------------------------------------------------
package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	for {
		now := time.Now()
		hour := now.Local().Hour()
		minutes := now.Local().Minute()
		seconds := now.Local().Second()

		clock := [...]placeholder{
			digits[hour/10],
			digits[hour%10],
			separator,
			digits[minutes/10],
			digits[minutes%10],
			separator,
			digits[seconds/10],
			digits[seconds%10],
		}

		screen.Clear()
		screen.MoveTopLeft()
		for line := range clock[0] {
			for i, digit := range clock {
				next := clock[i][line]
				if digit == separator && seconds%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}
