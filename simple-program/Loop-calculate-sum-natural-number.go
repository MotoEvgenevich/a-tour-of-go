// Sum of natural number (the positive numbers 1, 2, 3... are know as natural numbers)
package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Введите позитивное целое число : ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i // sum = sum + i
	}
	fmt.Print("Sum : ", sum)
}
