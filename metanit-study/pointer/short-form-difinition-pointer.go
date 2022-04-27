package main

import "fmt"

func main() {
	f := 2.3
	pf := &f

	fmt.Println("Address: ", pf)
	fmt.Println("Value: ", *pf)

	var pg *float64
	pg = pf
	if pg != nil {
		fmt.Println("Value: ", *pg)
	}

}
