package main

import "fmt"

func main() {
	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum int = 0
	for _, value := range numbers {
		if value < 0 {
			continue
		}
		sum += value
	}
	fmt.Println("Sum:", sum)
}
