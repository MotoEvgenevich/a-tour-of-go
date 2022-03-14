package main

import "fmt"

func main() {
	fmt.Print("Enter first string:")
	var first string
	fmt.Scanln(&first)
	fmt.Print("Enter two string:")
	var two string
	fmt.Scanln(&two)
	fmt.Print(first + two)
}

/* In this program, user is asked to enter two string.
The use of Print function and addition of two strings
without storing value in variable.

Strings and numbers are both extremely useful.
Concatenation uses the same symbol as addition.
The Go compiler figures out
what to do based on the types of the arguments.
Because both sides of the + are strings,
the compiler assumes you mean concatenation and
not addition (addition is meaningless for strings). */
