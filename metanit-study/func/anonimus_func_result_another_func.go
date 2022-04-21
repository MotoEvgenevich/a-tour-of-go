package main

import "fmt"

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }
	}
}

func main() {
	f := selectFn(1)
	fmt.Println(f(2, 3)) //5
	fmt.Println(f(4, 5)) //9
	fmt.Println(f(7, 6)) //13

	f = selectFn(2)
	fmt.Println(f(20, 3)) //17
	fmt.Println(f(4, 5))  //-1
	fmt.Println(f(7, 6))  //1

	f = selectFn(3)
	fmt.Println(f(2, 3)) //6
	fmt.Println(f(4, 5)) //20
	fmt.Println(f(7, 6)) //42
}

// Анонимная функция может быть результатом друой функции
