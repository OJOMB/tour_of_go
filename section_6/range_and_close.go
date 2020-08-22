package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func fibClosureToChannel(c chan int) func() {
	n0, n1 := 0, 1
	return func() {
		rv := n0
		n0, n1 = n1, n0+n1
		c <- rv
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	fib := fibClosureToChannel(c)

	wg.Add(2)
	go receiveAndPrint(c, &wg)
	go receiveAndPrint(c, &wg)

	for i := 0; i < 20; i++ {
		fib()
	}
	close(c)

	wg.Wait()

}

func receiveAndPrint(c <-chan int, wg *sync.WaitGroup) {
	id := RandString(5)
	i := 0
	for v := range c {
		fmt.Printf("%s %d: %d\n", id, i, v)
		i++
	}
	wg.Done()
}
