package main

import "fmt"

const kool = 3

func main() {
	fmt.Printf("type of kool: %T\n", kool)
	var u uint = 0
	var i uint = 1
	var sum uint = u - i

	fmt.Printf("integer has underflown: %d\n", sum) // integer underflow

	var u1 uint = 1<<64 - 1
	fmt.Printf("u: %d\n", u1)
	var i1 int = 1
	var sum1 uint = u1 + uint(i1)

	fmt.Printf("integer has overflown: %d\n", sum1) // integer overflow
}
