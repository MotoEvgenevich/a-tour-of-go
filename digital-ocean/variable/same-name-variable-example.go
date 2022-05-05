package main

import "fmt"

var num1 = 5

//num3 := 16        //!!!!!!!! Подходит только для объявления локальных переменных !!!!!!!! Если раскоменьтить expected declaration, found num3

func printNumbers() {
	num1 := 10
	num2 := 7

	fmt.Println(num1)
	fmt.Println(num2)
}

func main() {
	printNumbers()    // output : 10, 7
	fmt.Println(num1) // output :5
}
