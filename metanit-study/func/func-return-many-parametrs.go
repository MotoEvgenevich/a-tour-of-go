package main

import "fmt"

func main() {

	var age, name = add(4, 5, "Tom", "Simpson")
	fmt.Println(age)
	fmt.Println(name)
}

func add(x, y int, firstName, lastName string) (int, string) {
	var z = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

/* Функция add принимает четыре параметра: два числа и две строки.
Возвращает число (значение типа int) и строку.
Возвращаемые значения указаны после оператора return.

Поскольку функция теперь возвращает два значения,
то при вызове этой функции мы можем присвоить ее результат двум переменным:

var age, name = add(4, 5, "Tom", "Simpson")

Первое возвращаемое значение передается первой переменной age,
а второе значение передается второй переменной name.

Альтернативный способ передачи переменным результатов функции:

age, name := add(4, 5, "Tom", "Simpson")

И также в даном случае можно было бы использовать именованные результаты:

func add(x, y int, firstName, lastName string) (z int, fullName string) {
    z = x + y
    fullName = firstName + " " + lastName
    return
} */
