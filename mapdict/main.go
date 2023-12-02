package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Print("try again")
	}else{
		input := os.Args[1]

		dict := map[string]string{
			"cat":  "แมว",
			"dog":  "หมา",
			"cow":  "วัว",
			"duck": "เป็ด",
		}
	
		if value, ok := dict[input]; ok {
			fmt.Printf("%q แปลว่า %s", input, value)
		}
	}
}
