package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}
	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	separator := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	for {
		now := time.Now()
		hour := now.Local().Hour()
		minutes := now.Local().Minute()
		seconds := now.Local().Second()

		clock := [...]placeholder{
			digits[hour/10],
			digits[hour%10],
			separator,
			digits[minutes/10],
			digits[minutes%10],
			separator,
			digits[seconds/10],
			digits[seconds%10],
		}

		screen.Clear()
		screen.MoveTopLeft()
		for line := range clock[0] {
			for i, digit := range clock {
				next := clock[i][line]
				if digit == separator && seconds%2 == 0 {
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
