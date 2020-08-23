package main

import (
	"fmt"
	"math/rand"
	"time"
)

// messages sent to channel at random time intervals between 0 and 5 seconds
// the select statement will trigger return if the interval is greater than 4
// player wins if they randomly get n intervals less than 4 seconds

func gen(n int) chan time.Duration {
	var ch chan time.Duration = make(chan time.Duration)
	go func() {
		for i := 0; i < n; i++ {
			wait := time.Duration(rand.Float64()*5000) * time.Millisecond
			ch <- wait
			time.Sleep(wait)
		}
		close(ch)
	}()
	return ch
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	n := 10
	var ch chan time.Duration = gen(n)
	for i := 0; i < n; i++ {
		select {
		case msg := <-ch:
			fmt.Printf("level: %d, wait interval: %0.1fms\n", i, float64(msg/time.Millisecond))
		case <-time.After(4 * time.Second):
			fmt.Println("too slow friend, you blew it")
			return
		}
	}
	fmt.Println("oh boy you win!")
}
