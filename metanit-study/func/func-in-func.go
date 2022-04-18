package main

import "fmt"

func add(x int, y int) int {

	return x + y
}

func multiply(x int, y int) int {

	return x * y
}

func action(n1 int, n2 int, operation func(int, int) int) {

	result := operation(n1, n2)
	fmt.Println(result)
}

func main() {

	action(10, 25, add)    //35
	action(5, 6, multiply) //30
}

/* Здесь функция action принимает три параметра.
Первые два параметра - числа, а третий параметр - функция,
которая соответствует типу: func(int , int) int.
То есть третий параметр представляет некоторое действие и
может быть представлен любой функцией, которая принимает два значения типа int
и возвращает также значение типа int. Для примера здесь как раз определены две подобных функции,
которые соответствуют данному типу: add и multiply.
Через имя параметра operation мы сможем вызывать данную функцию. */
