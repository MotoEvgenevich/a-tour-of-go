package main

import "fmt"

var globalValue = 1

/* global variable, it is visible from everywrere in the package,
but you don't need use them, it's just an example */
func main() {
	var i = 2
	{
		var j = 3
		fmt.Println("внутренний скоуп")
		fmt.Println("globalValue =", globalValue)
		fmt.Println("i =", i)
		fmt.Println("j =", j)
	}
	fmt.Println("наружний скоуп")
	fmt.Println("globalValue =", globalValue)
	fmt.Println("i =", i)
	/* fmt.Println("j", j) - если раскоменнтировать, то будет
	ошибка, потому что j определен внутри своего скоупа
	и тут будет не виден */
}
