package main

import (
	"fmt"
	"math"
)

func compute(fn func(x, y float64) float64, x, y int) float64 {
	return fn(float64(x), float64(y))
}

func main() {

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(compute(hypot, 5, 64))

	fmt.Println(hypot(5, 64))
}
