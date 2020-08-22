package main

import (
	"fmt"
)

func main() {
	var c chan string = make(chan string, 5)
	c <- "hello"
	c <- "world"
	c <- "hello"
	c <- "world"
	c <- "hello"

	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
}
