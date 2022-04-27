package main

import "fmt"

func main() {
	var x int = 4
	var p *int = &x
	*p = 25
	fmt.Println(x)
}
