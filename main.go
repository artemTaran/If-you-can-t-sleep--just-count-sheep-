package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSheep(3))
}

func countSheep(num int) string {
	text := ""
	for i := 0; i < num; i++ {
		text += fmt.Sprintf("%d Sheep...", i+1)
	}
	return text
}
