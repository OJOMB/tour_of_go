package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c1 := make(chan int)
	c2 := make(chan int)

	c1Closed := make(chan int)
	c2Closed := make(chan int)

	go firstFive(c1, c2, c1Closed, c2Closed)

	start := time.Now().UTC()
	half := (1<<64 - 1) / 2
	for i := 0; true; i++ {
		if pick := rand.Int(); pick <= half {
			c1 <- i
		} else {
			c2 <- i
		}
		if _, open := <-c1Closed; !open {
			fmt.Println("channel one wins with a time of: %v\n", time.Now().UTC().Sub(start))
			return
		} else if _, open := <-c2Closed; !open {
			fmt.Println("channel two wins with a time of: %v\n", time.Now().UTC().Sub(start))
			return
		}
	}

}

func firstFive(c1 <-chan int, c2 <-chan int, c1Closed chan<- int, c2Closed chan<- int) {
	// the way this works is probably an anti-pattern but nevermind it's just for experimenting
	c1Count, c2Count := 0, 0
	for {
		select {
		case n := <-c1:
			fmt.Printf("received msg number %d from c1\n", n)
			c1Count += 1
		case n := <-c2:
			fmt.Printf("received msg number %d from c2\n", n)
			c2Count += 1
		}

		if c1Count == 5 {
			close(c1Closed)
			return
		} else if c2Count == 5 {
			close(c2Closed)
			return
		}

	}
}
