package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}
func isPositive(n int) bool {
	return n > 0
}

func sum(numbers []int, criteria func(int) bool) int {

	result := 0
	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}
	return result
}
func main() {

	slice := []int{-2, 4, 3, -1, 7, -4, 23}

	sumOfEvens := sum(slice, isEven) // сумма четных чисел
	fmt.Println(sumOfEvens)          // -2

	sumOfPositives := sum(slice, isPositive) // сумма положительных чисел
	fmt.Println(sumOfPositives)              // 37
}

/* Здесь функция sum вычисляет сумму элементов среза. Но н всех элементов, а только тех,
которые соответствуют условию. Условие передается в виде функции в качестве второго параметра.
Условие должно соответствовать функции типа func(int) bool. То есть функция должна принимать
в качестве параметра значение типа int и возвращать значение типа bool, которое указывает,
соответствовует ли переданное число условию.

Для примера здесь также определены две вспомогательные функции: isEven (возвращает true, если число четное)
и isPositive (возвращает true, если число положительное). Эти функции соответствуют типу func(int) bool,
поэтому их можно использовать в качестве условия. */
