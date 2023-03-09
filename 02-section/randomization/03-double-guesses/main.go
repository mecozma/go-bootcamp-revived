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
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
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
)

func main() {
	// seed the rand package with the time package.
	rand.Seed(time.Now().UnixNano())

	// get args from the terminal.
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println(usage)
		fmt.Println(example)
		return
	}

	guessOne, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("The argument provided is not a number.")
		return
	}

	guessTwo, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("The argument provided is not a number.")
		return
	}

	if guessOne < 1 || guessTwo < 1 {
		fmt.Println("Please provide a positive number.")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		r := float64(rand.Intn(int(math.Max(guessOne, guessTwo))) + 1)
		if r == guessOne || r == guessTwo {
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
