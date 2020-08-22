package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)
	for _, word := range strings.Split(s, " ") {
		_, ok := m[word]
		if ok {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}
	return
}

func main() {
	s := `
now what in the carpet world is going on here? 
I told you to have that carpet cleaned and dried 
by Wednesday and as far as I can tell that carpet 
has not even been beaten like a carpet should be beaten`
	m := WordCount(s)

	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
}
