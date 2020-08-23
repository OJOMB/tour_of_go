package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(name string, quit chan bool) chan string {
	var ch chan string = make(chan string)
	go func() {
	loop:
		for {
			select {
			case ch <- name:
				wait := time.Duration(rand.Float64()*1000) * time.Millisecond
				time.Sleep(wait)
			case <-quit:
				fmt.Println("received quit signal from main thread")
				break loop // break statements have goto like labels
			}
		}
		fmt.Println("closing channel: %s", name)
		close(ch)
		quit <- false
	}()
	return ch
}

func main() {
	var quit chan bool = make(chan bool)
	ch := boring("joe", quit)
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

	// here we're signalling to the goroutine at the other end of quit that it should close the ch channel and return
	quit <- true
	// we wait for a signal back to confirm that it's quit successfully
	<-quit

	_, ok := <-ch
	if !ok {
		fmt.Println("confirmed that channel is closed")
	}
}
