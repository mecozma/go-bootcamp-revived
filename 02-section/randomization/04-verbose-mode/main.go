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
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

var (
	usage            = "go run main.go [provide a number in the range 1 - 10]"
	example          = "go run main.go 7"
	maxTurns         = 5
	winMessageOne    = "You won!"
	winMessageTwo    = "It's your lucky day!"
	winMessageThree  = "You made a point!"
	bonusMessage     = "You rock it!"
	lostMessageOne   = "It is not your day!"
	lostMessageTwo   = "Try again!"
	lostMessageThree = "Maybe tomorrow!"
	lostMessageFour  = "That's all you got?"
	verbose          bool
)

func main() {
	// seed the rand package with the time package.
	rand.Seed(time.Now().UnixNano())

	// get args from the terminal.
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println(usage)
		fmt.Println(example)
		return
	}

	if args[0] == "-v" {
		verbose = true
	}

	// get the last argument in the arguments arr
	guessOne, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("The argument provided is not a number.")
		return
	}

	if guessOne < 1 {
		fmt.Println("Please provide a positive number.")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		r := rand.Intn(guessOne) + 1
		if verbose {
			fmt.Printf("%d ", r)
		}
		if r == guessOne {
			switch rand.Intn(3) {
			case 0:
				fmt.Println(winMessageOne)
			case 1:
				fmt.Println(winMessageTwo)
			case 2:
				fmt.Println(winMessageThree)
			}

			if turn == 1 {
				fmt.Println(bonusMessage)
			}
			return
		}
	}
	switch rand.Intn(5) {
	case 0:
		fmt.Println(lostMessageOne)
	case 1:
		fmt.Println(lostMessageTwo)
	case 2:
		fmt.Println(lostMessageThree)
	case 3:
		fmt.Println(lostMessageFour)
	}
}
