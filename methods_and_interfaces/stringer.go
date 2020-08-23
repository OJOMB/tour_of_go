package main

import "fmt"

// Stringer is one of the most commonly defined interfaces

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

	// the fmt package looks for whether a type implements Stringer when you pass an entity to on of it's Print functions
	// in this case if we comment out the String method we get this printed:
	// {Oscar unknown british 12345}
	// with the String method uncommented we get this printed instead:
	// <Person>name: Oscar, id: 12345</Person>
	fmt.Println(oscar)
}
