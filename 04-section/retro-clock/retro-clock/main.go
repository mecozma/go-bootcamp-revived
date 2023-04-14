// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Ticker: Slide the Clock
//
//	Your goal is slide the placeholders every second.
//	Please run the solution to see it in action.
//
//
//	THIS IS A HARD EXERCISE:
//	+ It will take you days but it will be worth it.
//	+ For experienced developers, this can take an hour or so.
//
//
//	1. You need to determine the starting and the ending digits to create
//	   the sliding effect.
//
//
//	2. Each second, start from the next placeholder, skip the previous one.
//	   This means: Only draw the next placeholders.
//
//	   Like this:
//
//	   12:40:31
//	   2:40:31
//	   40:31
//	   0:31
//	   :31
//	   31
//	   1
//
//
//	3. After the last placeholder is displayed, fill the lines for the missing
//	   placeholders, and then start from the first placeholder. Draw it to the
//	   right part of the screen.
//
//	   Like this:
//
//	   12:40:31
//	   2:40:31
//	   40:31
//	   0:31
//	   :31
//	   31
//	   1
//	          1
//	         12
//	        12:
//	       12:4
//	      12:40
//	     12:40:
//	    12:40:3
//	   12:40:31
//
//	   As you can see, you need to draw the clock from the right part of the
//	   screen, beginning from the first placeholder.
//
// HINTS
//   - You would need to clear the screen inside the loop instead of once.
//     Otherwise the previous placeholders will be left on the screen.
//
// EXPECTED OUTPUT
//
//	Please run the solution to see it in action. Do not look at the source-code
//	though.
//
// ---------------------------------------------------------
package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	for shift := 0; ; shift++ {

		screen.Clear()
		screen.MoveTopLeft()

		now := time.Now()
		hour := now.Local().Hour()
		minutes := now.Local().Minute()
		seconds := now.Local().Second()

		clock := [...]placeholder{
			digits[hour/10],
			digits[hour%10],
			colon,
			digits[minutes/10],
			digits[minutes%10],
			colon,
			digits[seconds/10],
			digits[seconds%10],
		}

		for line := range clock[0] {
			l := len(clock)
			// this sets the beginning and ending of the placeholder positions(indexes).
			// shift % 1 prevents the indexing error.
			s, e := shift%l, l

			// to slide the placeholders from the right part of the screen.
			//
			// here, we assume that as if the clock's length is double of its length.
			// this makes things easy to manage: that's why: l*2 is there.
			//
			// shift is always increasing, for it's to go beyond the clock's length,
			// it should be equal or greater than l*2, right (after the remainder of course)?
			//
			// so, if the clock goes beyond its length; this code detects that,
			// and resets the starting position to the first placeholder's index,
			// and it keeps doing so until the clock is fully displayed again.
			if shift%(l*2) >= l {
				s, e = 0, s
			}

			// print empty lines for the missing place holders.
			// this creates the effect of moving placeholders from right to left.
			//
			// l-e can only be non-zero when the above if statement runs.
			// otherwise, l-e is always zero, because l == e.
			//
			// this is one of the other benefits of assuming the length of the
			// clock as the double of its length. otherwise, l-e would always be 0.
			for j := 0; j < l-e; j++ {
				fmt.Print("     ")
			}

			// draw the digits starting from 's' to 'e'
			for i := s; i < e; i++ {
				next := clock[i][line]
				if clock[i] == colon && seconds%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}
