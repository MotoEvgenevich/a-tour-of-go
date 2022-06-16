package main

import "fmt"

func main() {
	var third []bool
	fmt.Printf("len(%d) cap(%d) = %v\n", len(third), cap(third), third)
}
