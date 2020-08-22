package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())
	// the go keyword indicates to the runtime that the function that follows
	// should be run in a user space thread handled by the go scheduler
	// a goroutine
	// goroutines are managed by the go scheduler
	go say("world")
	say("hello")

}

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(500 * time.Millisecond)
	}

}
