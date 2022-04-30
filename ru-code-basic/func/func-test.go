package main

import (
	"fmt"
	"strconv"
)

func main() {
	IntToString(-42)
	fmt.Println(IntToString(-42))
}

func IntToString(x int) string {

	s := strconv.Itoa(x)
	return s
}
