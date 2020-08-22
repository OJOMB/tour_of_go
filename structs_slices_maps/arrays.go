package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// what about multi-dimensional arrays?
	var multiD [2][5]float64 = [2][5]float64{
		{1.9, 2.8, 3.7, 4.6, 5.5},
		{2.4, 3.3, 4.2, 5.1, 6.0},
	}
	fmt.Println(multiD)
	fmt.Printf("multiD row 2, col 4: %0.1f\n", multiD[1][3])

	// In Golang, unlike C/C++/Java, arrays are value types.
	// this means that when an array is assigned to a new var, a whole new copy is made
	// so if you pass an array to a func or assign it to a new var you need to remember
	// you're not dealing with a reference to the original, you're dealing with a new copy
	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1 // A copy of the array `a1` is assigned to `a2`

	a2[1] = "German"

	fmt.Printf("a1[0]: %s\n", a1[0]) // a1 remains unchanged ("English")
	fmt.Printf("a2[0]: %s\n", a2[0]) // a2 is now "German"

}
