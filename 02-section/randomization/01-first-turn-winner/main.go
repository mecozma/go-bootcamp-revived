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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

var (
	usage        = "go run main.go [provide a number in the range 1 - 10]"
	example      = "go run main.go 7"
	maxTurns     = 5
	message      = "You won!"
	bonusMessage = "You rock it!"
	lostMessage  = "It is not your day!"
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
			fmt.Println(message)

			if turn == 1 {
				fmt.Println(bonusMessage)
			}
			return
		}
	}
	fmt.Println(lostMessage)
}
