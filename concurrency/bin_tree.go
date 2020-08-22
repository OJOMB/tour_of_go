package main

import "fmt"

type Tree struct {
	Left, Right *Tree
	Value       int
}

func walkTree(t *Tree) int {
	if t == nil {
		return 0
	}
	return t.Value + walkTree(t.Left) + walkTree(t.Right)
}

var kool Tree = Tree{
	Value: 3,
	Left: &Tree{
		Value: 1,
		Left: &Tree{
			Value: 1,
			Left:  nil,
			Right: nil,
		},
		Right: &Tree{
			Value: 2,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &Tree{
		Value: 8,
		Left: &Tree{
			Value: 5,
			Left:  nil,
			Right: nil,
		},
		Right: &Tree{
			Value: 13,
			Left:  nil,
			Right: nil,
		},
	},
}

func main() {
	fmt.Println(walkTree(&kool))
}
