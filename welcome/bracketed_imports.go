package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("the time is", time.Now())
	fmt.Println("the year is", time.Now().Year())
	fmt.Println("the month is", time.Now().Month())
	fmt.Println("the day is", time.Now().Day())
}
