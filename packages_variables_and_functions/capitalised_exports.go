package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		In Go, only capitalised objects are exported beyond their encapsulating package. Lowercase naming implies the
		object is private to the package. This is why we can use math.Pi but not math.mPi4 which is private
	*/

	// successful
	fmt.Println(math.Pi)
	// unsuccessful
	fmt.Println(math.mPi4)
}
