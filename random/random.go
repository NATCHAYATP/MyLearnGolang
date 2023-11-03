package main

import (
	"fmt"
	"math/rand"
)
// [1] basic random

// func main() {
// 	number := 10 // 0 - 10
// 	for i := 0; i != number; {
// 		i = rand.Intn(number + 1)
// 		fmt.Printf("%d ", i)
// 	}
// 	fmt.Println()
// }

// [2] basic random for use Seed
// กรณีอยากใช้ rand,Seed() ด้วย จะได้แบบนี้
// ถ้าคุณไม่ตั้งค่า seed ใหม่, และใช้ default seed ที่เดิมทุกครั้ง, อาจเกิดกรณีที่เลขสุ่มที่ได้จะซ้ำกันหรือไม่สุ่มเป็นแบบที่คุณต้องการได้ทุกครั้งที่คุณรันโปรแกรม.

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	number := 10 // 0 - 10
// 	for i := 0; i != number; {
// 		i = rand.Intn(number + 1)
// 		fmt.Printf("%d ", i)
// 	}
// 	fmt.Println()
// }

// [3] random easy game 
func main() {
	number := 10 // 0 - 10
	for i := 0; i<5; i++ {
		num := rand.Intn(number + 1)
		if num == number {
			fmt.Println("YOU WIN!")
			return
		}
	}
	fmt.Println("YOU LOST...")
}
