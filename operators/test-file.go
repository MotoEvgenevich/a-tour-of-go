package main

import (
	"fmt"
)

func main() {
	names()
}
func names() {
	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)
	// проверяем на глассные имя
	for i, v := range name {
		fmt.Println(i, v, string(v))
	}
}
