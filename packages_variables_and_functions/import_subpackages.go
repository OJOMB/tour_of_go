package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		setting the seed like this with some pseudorandom input (like the unix timestamp in nano seconds
		ensures that each of the calls to rand.Intn(n) will return different values each time the program is run.
		Otherwise rand.Seed defaults to 1 and we get the same output every time.
	*/
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("this is a random int:", rand.Intn(101))
	fmt.Println("this is a random int:", rand.Intn(304))
	fmt.Println("this is a random int:", rand.Intn(2060))
}
