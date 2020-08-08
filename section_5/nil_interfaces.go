package main

import "fmt"

type MyFloat float64

type Vector struct {
	X, Y float64
}

type Scalable interface {
	Scale(float64)
}

func main() {
	fmt.Println("============original vars=============")
	var v Vector = Vector{3, 4}
	var mf MyFloat = 42
	fmt.Println(v)
	fmt.Println(mf)

	fmt.Println("=============scaled vars==============")
	scaleByFive(&v)
	scaleByFive(&mf)

	fmt.Println(v)
	fmt.Println(mf)

	fmt.Println("=========using interface value=========")

	var i1 Scalable

	i1 = &v
	i1.Scale(2)

	i1 = &mf
	i1.Scale(2)

	fmt.Println(v)
	fmt.Println(mf)

	fmt.Println("=======using nil interface value========")

	var i2 Scalable
	var nill_v Vector
	var nill_mf MyFloat

	i2 = &nill_v
	i2.Scale(2)

	i2 = &nill_mf
	i2.Scale(2)

	fmt.Println(nill_v)
	fmt.Println(nill_mf)

}

func (v *Vector) Scale(sf float64) {
	if v == nil {
		fmt.Println("whadda flip? vector received nil")
	}
	v.X *= sf
	v.Y *= sf
}

func (f *MyFloat) Scale(sf float64) {
	if f == nil {
		fmt.Println("whomst the flip? myfloat received nil")
	}
	*f *= MyFloat(sf)
}

// function that takes interface type
func scaleByFive(scalable Scalable) {
	scalable.Scale(5)
}
