package main

import "fmt"

var g = "global"

func printLocal() {
	l := "local"
	fmt.Println(l)
}

func main() {
	//fmt.Println(l) 			Раскоменть чтобы увидеть ошибку компилятора Undeclared name in this scope!!!!
}
