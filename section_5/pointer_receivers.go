package main

import f "fmt"

type Vector struct {
	X, Y float64
}

func (v *Vector) Scale(sf int) {
	v.X *= float64(sf)
	v.Y *= float64(sf)
}

func (v Vector) Double() Vector {
	v.X *= float64(2)
	v.Y *= float64(2)
	return v
}

func main() {
	v := Vector{3, 5}
	v.Scale(5) // this modified v in place

	f.Println(v)
	f.Println(v.Double()) // this returned a new vector separate from v
	f.Println(v)

}
