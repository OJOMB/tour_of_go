package main

import "fmt"

func main() {
	// make can be used to create slic3es, maps and channels\
	// make takes a type as it's first arg and returns a value of the same type
	// in the case of slice, the next 2 args are length and capacity
	// make([]int, 0, 10)
	// this creates an underlying array of length 10 and slice of that array of length 0
	// naturally capacity must be at least as large as length
	a := []int{1, 2}
	b := [2]int{11, 22}
	a = append(a, b...)

	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
