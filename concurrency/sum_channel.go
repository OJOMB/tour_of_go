package main

import (
	"fmt"
	"sync"
)

func sum(s []int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	var wg sync.WaitGroup
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	c := make(chan int)

	start := 0
	for i := 0; i < 2; i++ {
		go sum(s[start:start+10], c, &wg)
		wg.Add(1)
		start += 10
	}
	x, y := <-c, <-c // receive from c
	wg.Wait()
	fmt.Println(x, y)
}
