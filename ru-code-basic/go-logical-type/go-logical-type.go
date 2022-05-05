package main

import "fmt"

func main() {
	IsValid(10, "hello world")
	fmt.Println(IsValid(20, "hello"))
}
func IsValid(id int, text string) bool {
	var x int = id
	x = 10
	var y string = ""
	y = "hello"
	var isTrue bool = true
	var isFalse bool = false
	if x > 0 && y != "" {
		return isTrue
	} else {
		return isFalse
	}

}

/*
IsValid(0, "hello world") // false
IsValid(-22, "hello world") // false
IsValid(22, "") // false
IsValid(225, "hello world") // true
*/

/*
Реализуйте функцию IsValid(id int, text string) bool,
которая проверяет, что переданный идентификатор
id является положительным числом и текст text не пустой.
*/
