package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

type Person struct {
	first_name  string
	surname     string
	gender      string
	eye_colour  string
	hair_colour string
	height      float64
	hobbies     []string
}

func detailPerson(p *Person) {
	var pronoun string
	switch p.gender {
	case "male":
		pronoun = "he has"
	case "female":
		pronoun = "she has"
	default:
		pronoun = "they have"
	}

	fmt.Printf(
		"%s %s has %s eyes and %s hair and is %0.1fcms tall. %s the following hobbies: ",
		p.first_name, p.surname, p.eye_colour, p.hair_colour, p.height, strings.ToUpper(pronoun)
	)
	fmt.Printf("%s.\n", strings.Join(p.hobbies, ", "))
}

func main() {
	var v Vertex = Vertex{20, 1}

	fmt.Printf("vertex v: %v\n", v)

	fmt.Printf("v.x: %d\n", v.X)
	fmt.Printf("v.y: %d\n", v.Y)

	var oscar Person = Person{
		first_name:  "Oscar",
		surname:     "Bardrick",
		gender:      "male",
		eye_colour:  "blue",
		hair_colour: "brown",
		height:      175.0,
		hobbies: []string{
			"programming",
			"cinema",
			"chilling in the park",
		},
	}

	detailPerson(&oscar)
}
