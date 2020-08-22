package main

import (
	"fmt"
	"math/cmplx"
)

// declare multiple vars like so:
var (
	ToBe    bool       = true
	NotToBe bool       = false
	MaxInt  uint64     = 1<<64 - 1
	z       complex128 = -5 + 12i
	rootz   complex128 = cmplx.Sqrt(z)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", rootz, rootz)
}
