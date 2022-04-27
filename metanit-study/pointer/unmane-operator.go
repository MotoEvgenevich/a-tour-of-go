package main

import "fmt"

func main() {
	var x int = 4
	var p *int = &x
	fmt.Println("Address: ", p)
	fmt.Println("Value: ", *p)
}
