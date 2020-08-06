package main

import "fmt"

var i int = 1
var j, k int = 1, 2

// if an initialisation value is present we can omit the type sig and Go will infer it

var initialised = 100

func main() {
	// we can use a walrus guy to declare new vars without the var keyword as a shorthand
	c, python, java := true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
