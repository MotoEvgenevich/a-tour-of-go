package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(400 * time.Millisecond)
LOOP:
	for {
		select {
		case <-tick:
			i := 0
			fmt.Println("tick", i)
			break LOOP
		default:
			// "i" is undefind here
			fmt.Println("sleep")
			time.Sleep(30 * time.Millisecond)
		}
	}
	// "i" is undefined here
}
