package main

import "fmt"

type MyFloat float64

type Vector struct {
	x, y float64
}

// here we've defined an interface with a method that returns an empty interface
// this means that as long as a type T has a method called Scale that takes a float it implements the interface
// it doesn't matter what the type of the return value is since interface{} can hold anything
type Scaler interface {
	Scale(scaleFactor float64) interface{}
}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	fived := ScaleByFive(MyFloat(4))

	fmt.Println(fived)
}

func describe(i interface{}) {
	fmt.Printf("(value: %v, type: %T)\n", i, i)
}

// defining Scale for both MyFloat and Vector so they satisfy the interface
func (mf MyFloat) Scale(scaleFactor float64) interface{} {
	return MyFloat(scaleFactor) * mf
}

func (v Vector) Scale(scaleFactor float64) interface{} {
	v.x *= scaleFactor
	v.y *= scaleFactor
	return v
}

func ScaleByFive(scaler Scaler) interface{} {
	return scaler.Scale(5)
}
