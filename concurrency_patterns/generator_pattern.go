package main

import (
	"fmt"
	"math/rand"
	"time"
)

// we're using a generator function to generate a channel
// then we're multiplexing channels into a single master
// probably the same effect coulf be achieved with a for select

func generator(name string, wait time.Duration) chan string {
	var ch chan string = make(chan string)
	go func(name string, c chan string, wait time.Duration) {
		for {
			c <- name
			time.Sleep(wait)
		}
	}(name, ch, wait)
	fmt.Printf("generated channel with msg: %s, with wait duration %d\n", name, wait)

	return ch
}

func chmux(ch1, ch2 chan string) chan string {
	var ch = make(chan string)
	go func() {
		for {
			select {
			case msg := <-ch1:
				ch <- msg
			case msg := <-ch2:
				ch <- msg
			}
		}
	}()

	return ch
}

func main() {
	rand.Seed(time.Now().UnixNano())
	rando := func() time.Duration { return time.Duration(rand.Intn(5)) * time.Second }
	m1 := chmux(generator("o", rando()), generator("s", rando()))
	m2 := chmux(generator("c", rando()), generator("a", rando()))
	m3 := chmux(m1, m2)
	m4 := chmux(m3, generator("r", rando()))

	for msg := range m4 {
		fmt.Print(msg)
	}
}
