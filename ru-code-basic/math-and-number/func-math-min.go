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

func MinInt(x, y float64) float64 {
	var z float64 = math.Min(x, y)
	return z

}
