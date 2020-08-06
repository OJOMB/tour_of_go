package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x == 1 {
		return x
	}
	var acceptable_error float64 = 0.000000000000001
	var z float64 = 1

	for {
		old_z := z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("z: %0.20f\n", z)
		if math.Abs(old_z-z) < acceptable_error {
			return z
		}
	}
}

func main() {
	sqroot := Sqrt(35267)
	fmt.Printf("final z: %0.20f\n", sqroot)
	fmt.Printf("final z squared: %0.20f\n", math.Pow(sqroot, 2))
}
