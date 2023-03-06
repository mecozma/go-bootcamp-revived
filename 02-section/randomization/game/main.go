package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game!
	The program will pick a %d random numbers.
	Youd task is to buess one of those numbers.
	
	The greater your number is, the harder it gets.fmt
	
	Pick a number!?!`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Only pick positive numbers.")
		return
	}

	guess1, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Not a number")
	}

	if guess1 < 0 {
		fmt.Println("Only pick positive numbers.")
		return
	}
	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		if n == guess || n == guess1 {
			fmt.Printf("It's a win %d ", n)
			if turn == 0 {
				fmt.Println("... on the first try :)")
			}
			return
		}

		switch rand.Intn(4) {
		case 0:
			fmt.Println("Oh no! you lost")
			return
		case 1:
			fmt.Println("You didn't win!")
			return
		case 2:
			fmt.Println("You lost... Try again!")
			return

		}
	}
}
