package main

import "fmt"

type Vector struct {
	x, y float64
}

func main() {
	var kool []interface{} = make([]interface{}, 5, 10)
	kool = append(kool, 12, 69.76, "hello")
	fmt.Println(kool)
}
