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
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Moodly
//
//   1. Get username from command-line
//
//   2. Display the usage if the username is missing
//
//   3. Create an array
//     1. Add three positive mood messages
//     2. Add three negative mood messages
//
//   4. Randomly select and print one of the mood messages
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name]
//
//   go run main.go Socrates
//     Socrates feels good 👍
//
//   go run main.go Socrates
//     Socrates feels bad 👎
//
//   go run main.go Socrates
//     Socrates feels sad 😞
//
//   go run main.go Socrates
//     Socrates feels happy 😀
//
//   go run main.go Socrates
//     Socrates feels awesome 😎
//
//   go run main.go Socrates
//     Socrates feels terrible 😩
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Ussage:")
		fmt.Println("go run main.go [username][positive/negative]")
		return
	}

	moods := [...][3]string{
		{"happy 😀",
			"awesome 😎",
			"good 👍",
		},
		{"bad 👎",
			"sad 😞",
			"terrible 😩",
		},
	}
	mood := ""

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(3)

	switch args[1] {
	case "positive":
		mood = moods[0][r]
	case "negative":
		mood = moods[1][r]
	}

	// var mi int
	// name, mood := args[0], args[1]
	// if mood != "positive" {
	// 	mi = 1
	// }

	// fmt.Printf("%s feels %s\n", name, moods[mi][r])

	fmt.Printf("%s feels %s\n", args[0], mood)
}
