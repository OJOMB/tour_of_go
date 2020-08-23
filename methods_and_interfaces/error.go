package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf(
		"at %v, %s",
		e.When,
		e.What, // i just realised you can add a newline with a comma
	)
}

func getTimeIfEqualSeconds() (time.Time, error) {
	if now := time.Now(); now.Second()%2 == 0 {
		return now, nil
	}
	return nil, &MyError{time.Now(), "it didn't work"}
}

func main() {

	if now, err := run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(now)
	}
}
