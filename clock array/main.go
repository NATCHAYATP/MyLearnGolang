package main

import (
	"fmt"
	"time"
	"github.com/inancgumus/screen"
)

func main() {
	type pleaceholder [5]string
	zero := pleaceholder{
		"###",
		"# #",
		"# #",
		"# #",
		"###",
	}

	one := pleaceholder{
		"## ",
		" # ",
		" # ",
		" # ",
		"###",
	}

	two := pleaceholder{
		"###",
		"  #",
		"###",
		"#  ",
		"###",
	}

	three := pleaceholder{
		"###",
		"  #",
		"###",
		"  #",
		"###",
	}

	four := pleaceholder{
		"# #",
		"# #",
		"###",
		"  #",
		"  #",
	}

	five := pleaceholder{
		"###",
		"#  ",
		"###",
		"  #",
		"###",
	}

	six := pleaceholder{
		"###",
		"#  ",
		"###",
		"# #",
		"###",
	}

	seven := pleaceholder{
		"###",
		"  #",
		"  #",
		"  #",
		"  #",
	}

	eight := pleaceholder{
		"###",
		"# #",
		"###",
		"# #",
		"###",
	}

	nine := pleaceholder{
		"###",
		"# #",
		"###",
		"  #",
		"###",
	}

	colon := pleaceholder{
		"   ",
		" # ",
		"   ",
		" # ",
		"   ",
	}
	digits := [...]pleaceholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	screen.Clear()
	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		// [8][5]string --> [8]peachholder
		clock := [...]pleaceholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2==0{
					next = "   "
				}
				fmt.Print(next, " ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}

	// for line := range digits[0] {
	// 	for digit := range digits {
	// 		fmt.Print(digits[digit][line], " ")
	// 	}
	// 	fmt.Println()
	// }
}
