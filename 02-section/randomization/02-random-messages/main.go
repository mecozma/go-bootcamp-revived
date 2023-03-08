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
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
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
	if len(args) != 1 {
		fmt.Println(usage)
		fmt.Println(example)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("The argument provided is not a number.")
		return
	}

	if guess < 1 {
		fmt.Println("Please provide a positive number.")
	}

	for turn := 1; turn <= maxTurns; turn++ {
		r := rand.Intn(guess) + 1
		if r == guess {
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
