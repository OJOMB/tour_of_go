package main

import "fmt"

func main() {
	var pi *int
	var i int = 100
	pi = &i

	fmt.Printf("value of pi: %v\n", pi)
	fmt.Printf("dereferenced pi: %d\n", *pi)

	// we can manipulate the memory at the address given by the value of pi
	*pi /= 20
	fmt.Println(i)

	var nullPointer *float64

	fmt.Println(*nullPointer)

}
