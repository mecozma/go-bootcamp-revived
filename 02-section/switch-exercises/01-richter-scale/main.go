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
	"strconv"
)

// ---------------------------------------------------------
// STORY
//
//	You're curious about the richter scales. When reporters
//	say: "There's been a 5.5 magnitude earthquake",
//
//	You want to know what that means.
//
//	So, you've decided to write a program to do that for you.
//
// EXERCISE: Richter Scale
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//
//	go run main.go
//	  Give me the magnitude of the earthquake.
//
//	go run main.go ABC
//	  I couldn't get that, sorry.
//
//	go run main.go 0.5
//	  0.50 is micro
//
//	go run main.go 2.5
//	  2.50 is very minor
//
//	go run main.go 3
//	  3.00 is minor
//
//	go run main.go 4.5
//	  4.50 is light
//
//	go run main.go 5
//	  5.00 is moderate
//
//	go run main.go 6
//	  6.00 is strong
//
//	go run main.go 7
//	  7.00 is major
//
//	go run main.go 8
//	  8.00 is great
//
//	go run main.go 11
//	  11.00 is massive
//
// ---------------------------------------------------------
var (
	no_arg      = "Give me the magnitude of the earthquake."
	no_nu       = "I couldn't get that, sorry."
	description string
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println(no_arg)
		return
	}

	grade, err := strconv.ParseFloat(arg[0], 64)
	if err != nil {
		fmt.Println(no_nu)
		return
	}

	switch g := grade; {
	case g < 2:
		description = "micro"
	case g >= 2 && g < 3:
		description = "very minor"
	case g >= 3 && g < 4:
		description = "minor"
	case g >= 4 && g < 5:
		description = "light"
	case g >= 5 && g < 6:
		description = "moderate"
	case g >= 6 && g < 7:
		description = "strong"
	case g >= 7 && g < 8:
		description = "major"
	case g >= 8 && g < 11:
		description = "great"
	default:
		description = "massive"
	}
	fmt.Printf("%.2f is %s\n", grade, description)
}
