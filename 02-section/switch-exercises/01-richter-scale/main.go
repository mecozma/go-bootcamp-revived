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
	no_arg = "Give me the magnitude of the earthquake."
	no_nu  = "I couldn't get that, sorry."
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

	switch grade {
	case 0.5:
		fmt.Printf("%.2f is micro\n", grade)
	case 2.5:
		fmt.Printf("%.2f is very minor\n", grade)
	case 3:
		fmt.Printf("%.2f is minor\n", grade)
	case 4.5:
		fmt.Printf("%.2f is light\n", grade)
	case 5:
		fmt.Printf("%.2f is moderate\n", grade)
	case 6:
		fmt.Printf("%.2f is strong\n", grade)
	case 7:
		fmt.Printf("%.2f is major\n", grade)
	case 8:
		fmt.Printf("%.2f is great\n", grade)
	case 11:
		fmt.Printf("%.2f is massive\n", grade)
	}
}
