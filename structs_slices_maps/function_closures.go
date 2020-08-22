package main

import f "fmt" // realised you can alias imports like this

func adder() func(int) int {
	sum := 0 // setup free var which inner func has access to
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		f.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
