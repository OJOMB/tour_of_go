package main

import (
	"fmt"
)

func main() {
	var s []uint8
	s = append(s, 'a')
	s = append(s, 255)

	fmt.Println(s)
	fmt.Printf("s is of type: %T\n", s)

	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var sliceOfA []int = a[3:7]

	fmt.Println(sliceOfA)
	fmt.Printf("sliceOfA is of type: %T\n", sliceOfA)

	sliceOfA = append(sliceOfA, 899)
	fmt.Println(sliceOfA)
	fmt.Println(a)
}
