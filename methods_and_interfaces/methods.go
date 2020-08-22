package main

import (
	"fmt"
	"math"
)

// we can write methods for non-struct types too
type MyFloat float64

type Vertex struct {
	X, Y float64
}

// methods are functions with receiver arguments
// receiver args are specified after the func keyword and before the function name
// ULTIMATELY METHODS ARE JUST FUNCTIONS
// we can write a method as function by placing the receiver arg in the normal arg list
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return MyFloat(-f)
	}
	return f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	var f MyFloat = 69 - 420
	fmt.Printf("Absolute value of %0.02f is %0.02f\n", f, f.Abs())
}
