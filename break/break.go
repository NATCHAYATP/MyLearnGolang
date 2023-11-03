package main

import "fmt"

func main() {
	// use break all loop
	test:
	for i:=0; i<3;i++ {
		for {
			fmt.Print("break\n")
			break test
		}
		//fmt.Printf("continue %d\n",i)
	}
}
