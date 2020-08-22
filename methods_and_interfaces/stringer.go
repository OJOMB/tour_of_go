package main

import "fmt"

// Stringer is one of the most commonly defined interfaces
// the fmt package looks for it when you pass an entity to a Print function and calls it if implemented

type Person struct {
	name        string
	dob         string
	nationality string
	id          uint
}

func (p Person) String() string {
	return fmt.Sprintf("<Person>name: %s, id: %d</Person>", p.name, p.id)
}

func main() {
	var oscar Person = Person{
		name:        "Oscar",
		dob:         "unknown",
		nationality: "british",
		id:          12345,
	}

	// "<Person>name: Oscar, id: 12345</Person>"
	fmt.Println(oscar)
}
