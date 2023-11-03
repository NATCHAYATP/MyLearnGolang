package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"

)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}
	// // เราลองเขียน
	// const pageSize = 4
	// for from := 0; from < len(items); from += pageSize{
	// 	to := from + pageSize
	// 	if to > len(items) {
	// 		to = len(items)
	// 	}
	// 	fmt.Print(items[from:to])
	// }

	const pageSize = 4

	l := len(items)
	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}
		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d", (from / pageSize))
		s.Show(head, currentPage)
	}
}
