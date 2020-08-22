package main

import f "fmt"

func fibClosure() func() int {
	n1 := 1
	n0 := 0
	return func() int {
		rv := n0
		n0, n1 = n1, n0+n1
		return rv
	}
}

func getFib(n int) int {
	fib := fibClosure()
	for i := 0; i < n; i++ {
		fib()
	}
	return fib()
}

func recursiveFib(n int) int {
	if n <= 1 {
		return n
	}
	return (recursiveFib(n-1) + recursiveFib(n-2))
}

func main() {
	fib := fibClosure()

	for i := 0; i < 10; i++ {
		f.Println(fib())
	}

	f.Println("-------getFib-------")

	for i := 0; i < 10; i++ {
		f.Println(getFib(i))
	}

	f.Println("-------recursiveFib-------")
	for i := 0; i < 10; i++ {
		f.Println(recursiveFib(i))
	}
}
