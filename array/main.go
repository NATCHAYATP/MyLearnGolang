package main

import (
	"fmt"
)

// func main() {
// 	moods := []string{"sad", "smile", "cry", "happy", "ok"}
	
// 	// name := os.Args[1] แบบนี้จะไม่ return แต่จะขึ้น error
// 	name := os.Args[1:]
// 	if len(name) != 1 {
// 		fmt.Print("input name")
// 		return
// 	}

// 	mood := rand.Intn(len(moods))

// 	fmt.Println(name, " is ", moods[mood])
// }

// func main() {
// 	std := [...][3]float64{
// 		{5, 6, 1},
// 		{9, 8, 4},
// 	}
// 	var sum float64
// 	for _, j := range std {
// 		for _, i := range j {
// 			sum += i
// 		}
// 	}
// 	num := len(std)*len(std[1])
// 	fmt.Println("sum = ", sum)
// 	fmt.Println("num = ", num)
// 	fmt.Print("AVG = ", sum/float64(num))
// }

func main() {
	names := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "emc2"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	for _, i := range names {
		for _, j := range i {
			fmt.Printf("%2s ",j)
		}
		fmt.Println()
	}
}

