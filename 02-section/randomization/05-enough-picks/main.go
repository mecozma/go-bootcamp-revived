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
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
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
	r                int
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
		if guessOne <= 10 {
			r = rand.Intn(10)
		} else {
			r = rand.Intn(guessOne) + 1
		}

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
