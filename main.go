package main

import (
	"fmt"
	"time"
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
	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()
	for {

		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}
		// fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, " ")
			}
			fmt.Println()
		}
		fmt.Println()

		time.Sleep(time.Second)
	}
}

// for _, digit := range digits {
// 	for _, line := range digit {
// 		fmt.Println(line)
// 	}

// 	fmt.Println()
// }

// type (
// 	integer  int
// 	cabinet  [5]int
// 	bookcase [5]integer
// )

// _ = [5]integer{} == cabinet{}

// blue := bookcase{6, 9, 3, 2, 1}
// red := cabinet{6, 9, 3, 2, 1}

// fmt.Print("Are they equal? ")

// if cabinet(blue) == red {
// 	fmt.Println("YES")

// } else {
// 	fmt.Println("NO")
// }

// fmt.Printf("blue: %#v\n", blue)
// fmt.Printf("red: %#v\n", bookcase(red))
