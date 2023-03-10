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
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game too easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
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
	dynamicDificulty int
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

	if guessOne > 10 {
		dynamicDificulty = guessOne / 4
	}

	for turn := 1; turn <= maxTurns+dynamicDificulty; turn++ {
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
