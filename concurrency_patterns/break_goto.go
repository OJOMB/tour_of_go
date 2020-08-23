package main

import "fmt"

// go has goto like c
// break statements use goto like labels
// break labels identify the block of code to break free from

func main() {
	fmt.Println("no escape")
loop:
	for i := 0; ; i++ {
		if i >= 10 {
			break loop
		}
	}
	goto loop
}
