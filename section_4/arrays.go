package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// what about multi-dimensional arrays
	var multiD [2][5]float64 = [2][5]float64{
		{1.9, 2.8, 3.7, 4.6, 5.5},
		{2.4, 3.3, 4.2, 5.1, 6.0},
	}

	fmt.Printf("multiD row 2, col 4: %0.1f\n", multiD[1][3])
}
