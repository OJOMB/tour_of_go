package main

import "fmt"

// super simple interface
type Greeter interface {
	hello() string
}

type Person struct {
	name        string
	dateOfBirth string
	greeting    string
}

type Animal struct {
	name  string
	legs  int
	tail  bool
	noise string
}

func main() {
	var cow Animal = Animal{
		name:  "cow",
		legs:  4,
		tail:  true,
		noise: "moooo",
	}

	var oscar Person = Person{
		name:        "Oscar",
		dateOfBirth: "unknown",
		greeting:    "alright?",
	}

	// interface value, not sure what the point of this is
	var g Greeter

	g = oscar

	fmt.Println(g.hello())

	g = &cow // has to be apointer since greeter is implemented by *Animal

	fmt.Println(g.hello())

	// generic function
	Kool(oscar)
	Kool(&cow)
}

// *Animal and Person implement the interface Greeter

func (person Person) hello() string {
	if person.greeting != "" {
		return person.greeting
	}
	return "I say, nice to meet you!"
}

func (animal *Animal) hello() string {
	if animal.noise != "" {
		return animal.noise
	}
	return "..."
}

// function that takes an interface and is therefore generic over all types that implement that interface
func Kool(g Greeter) {
	fmt.Println(g.hello())
}
