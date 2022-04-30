package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		a = 7
		b = 4
	)
	MinInt(a, b)
	fmt.Println(MinInt(a, b))

}

func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}
