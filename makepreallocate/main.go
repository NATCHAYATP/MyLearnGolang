package main

import (
	"strings"

	s "github.com/inancgumus/prettyslice"
)

//วิธีที่ 1 ptr จะเปลี่ยนตลอดเพราะ cap มันไม่พอ 
// func main() {
// 	text := []string{"cat", "dog", "fish"}
// 	var upperText []string

// 	s.Show("before ", text)

// 	for _, str := range text {
// 		upperText = append(upperText, strings.ToUpper(str))
// 		s.Show("after ", upperText)
// 	}

// 	s.Show("after ", upperText)
// }

// วิธีที่ 2 จะทำให้ ptr เป็นตัวเดิม โดยใช้ make มาช่วย และจองพื้นที่ 
// จะเห็นว่ามันไม่ไปใส่ในพื้นที่ที่จองไว้ ถึง ptr จะเหมือนกันแต่ก็ต้องปรับ ที่ไม่เอาไปใส่เพราะมันมีค่าเป็น 0 อยู่
// func main() {
// 	text := []string{"cat", "dog", "fish"}
// 	// len = 0 , capacity = len(text)
// 	upperText := make([]string, len(text))

// 	s.Show("before ", text)

// 	for _, str := range text {
// 		upperText = append(upperText, strings.ToUpper(str))
// 		s.Show("after ", upperText)
// 	}

// 	s.Show("after ", upperText)
// }

// วิธีที่ 2 แบบปรับ คือให้จองพื้นที่แบบความจุ จะได้ไม่มีค่าอยู่
func main() {
	text := []string{"cat", "dog", "fish"}
	// len = 0 , capacity = len(text)
	upperText := make([]string, 0, len(text))

	s.Show("before ", text)

	for _, str := range text {
		upperText = append(upperText, strings.ToUpper(str))
		s.Show("after ", upperText)
	}

	s.Show("after ", upperText)
}