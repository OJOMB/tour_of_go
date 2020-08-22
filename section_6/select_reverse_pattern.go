package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func firstToN(n int, ch1 chan int, ch2 chan int) {
	rand.Seed(time.Now().UnixNano())
	ch1count, ch2count := 0, 0
	for i := 0; true; i++ {
		if randomNum := rand.Float64(); randomNum <= 0.5 {
			ch1 <- i
			ch1count++
		} else {
			ch2 <- i
			ch2count++
		}
		if ch1count >= n {
			close(ch1)
			return
		} else if ch2count >= n {
			close(ch2)
			return
		}
	}
}

func main() {
	var ch1, ch2 chan int = make(chan int), make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go firstToN(10, ch1, ch2)

	for {
		select {
		case msg, ok := <-ch1:
			if !ok {
				fmt.Println("channel one is the victor\n")
				return
			}
			fmt.Printf("channel one received: %d\n", msg)
		case msg, ok := <-ch2:
			if !ok {
				fmt.Println("channel two is the victor\n")
				return
			}
			fmt.Printf("channel two received: %d\n", msg)
		}

	}
}
