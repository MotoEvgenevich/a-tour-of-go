package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var f func(int, int) int = add
	fmt.Println(f(3, 4))

	var x = f(4, 5) //9
	fmt.Println(x)

}

/* Здесь переменная f имеет тип func(int, int) int, то есть представляет любую функцию,
которая принимает два параметра типа int и возвращает значение типа int.
Поэтому мы можем присвоить этой переменной функцию add,
которая соответствует данному типу:

var f func(int, int) int = add

После этого мы можем вызывать присвоенную функцию через переменную,
передавая нужные значения для ее параметров:

var x = f(4, 5) // 9 */
