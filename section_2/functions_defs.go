package main

import "fmt"

/*
	Functions are declared with the func keyword.
	Type signatures are described to the right of the var name.

	Return type is given after the argument brackets:
		func kool(x T1, y T2) T3 {}

	You can return multiple values from a function:
		func kooler(x T1, y T2) (T3, T4) {}

	if T1 and T2 are the same we can shorten things up like so:
		func koolerer(x, y T1) (T3, T4) {}

	you can also name the return values in the function signature:
		func koolest(x T1, y T2) (a, b T3) {
			a = "hmmm"
			b = "hwhat"
			return
		}
*/

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func swap1(x, y string) (a, b string) {
	a = y
	b = x
	// naked return
	return
}

func main() {
	fmt.Println("add(42, 13):", add(42, 13))
	fmt.Println("add1(42, 13):", add1(42, 13))
	a, b := swap1("x", "y")
	fmt.Printf("swap('x', 'y'): x = %s, y = %s\n", a, b)
	a, b = swap1("x1", "y1")
	fmt.Printf("swap1('x1', 'y1'): x1 = %s, y1 = %s\n", a, b)
}
