package main

import "fmt"

func main() {
	test()

}

func test() {
	var balance []int = []int{1, 2, 3, 6, 9, 12, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 36, 42, 90}

	for _, num := range balance {
		numb_test := num % 3
		fmt.Println(num)
	}
	if numb_test == 0 {
		fmt.Printf("number: %d ", numb_test)
		return
	}

}
