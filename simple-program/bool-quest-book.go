package main

import "fmt"

// Задание : Какое значение примет выражение (true && false) || (false && true) || !(false && false)?

func main() {
	var a bool = true
	var b bool = false
	var c bool = false
	var d bool = true
	var e bool = false
	var f bool = false

	var x bool = (a && b) || (c && d) || !(e && f)
	{
		fmt.Println(x)
	}
}
