package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)

	// sort แค่ ตรงกลาง
	mid := len(items) / 2
	smid := items[mid-1 : mid+2]

	// sort ปกติก็แค่เอา items มาใส่แทน smid
	sort.Strings(smid)

	fmt.Println()
	fmt.Println("Sorted  :", items)
}