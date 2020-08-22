package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = 4
	var f float64 = math.Sqrt(float64(x))
	var z uint = uint(f)
	fmt.Printf("x -> Type: %T, value: %d\nf -> Type: %T, value: %f\nz ->  Type: %T, value: %d\n", x, x, f, f, z, z)
}
