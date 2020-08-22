package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

type Animal struct {
	name    string
	legs    int
	tail    bool
	eyes    int
	fur     bool
	genus   string
	goodBoy bool
}

var (
	m map[string]Animal
)

func main() {
	m = make(map[string]Animal)
	m["cat"] = Animal{
		name:    "cat",
		legs:    4,
		tail:    true,
		eyes:    2,
		fur:     true,
		genus:   "felis",
		goodBoy: false,
	}
	m["dog"] = Animal{
		name:    "dog",
		legs:    4,
		tail:    true,
		eyes:    2,
		fur:     true,
		genus:   "canis",
		goodBoy: true,
	}
	m["snek"] = Animal{
		name:    "snek",
		legs:    0,
		tail:    true,
		eyes:    2,
		fur:     false,
		genus:   "serpens",
		goodBoy: false,
	}

	fmt.Printf("CAT: %v\n", m["cat"])
	fmt.Printf("DOG: %v\n", m["dog"])

	// if we create a map initialised to a map literal we can omit the value type
	var m1 map[string]Vertex = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Italska":   {50.077874, 14.4361127},
	}

	fmt.Printf("m1 before delete: %v\n", m1)

	delete(m1, "Bell Labs")

	fmt.Printf("m1 after delete: %v\n", m1)

	elem, ok := m1["Bell Labs"]

	fmt.Println(ok)
	fmt.Println(elem)

	elem, ok = m1["Italska"]

	fmt.Println(ok)
	fmt.Println(elem)

}
