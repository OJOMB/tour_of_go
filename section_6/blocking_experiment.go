package main

import "fmt"

func main() {
	var ch chan int = make(chan int, 2)
	close(ch)
	v, open := <-ch
	fmt.Println("we didnt get blocked boys")
	fmt.Printf("value; %v,  open: %v", v, open)
}
