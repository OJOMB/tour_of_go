package main

import "fmt"

func fizzbuzz(min int, max int) {
	for i := min; i < max; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else if i%3 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz(1, 16)

	var kool [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, item := range kool {
		fmt.Printf("index: %d, item %d\n", index, item)
	}

	var n int = 100

	for n > 5 {
		fmt.Printf("n: %d\n", n)
		n -= 5
	}

	for true {
		fmt.Println("This will loop to infinity")
	}

}
