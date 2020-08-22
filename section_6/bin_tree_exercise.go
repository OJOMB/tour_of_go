package main

import (
	"fmt"
	"sync"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var (
		b        bool = false
		s1, s2   []int
		ch1, ch2 chan int = make(chan int), make(chan int)
	)

	go func() { Walk(t1, ch1); close(ch1) }()
	for i := range ch1 {
		s1 = append(s1, i)
	}
	go func() { Walk(t2, ch2); close(ch2) }()
	for j := range ch2 {
		s2 = append(s2, j)
	}
	if len(s1) != len(s2) {
		return false
	}

	fmt.Println("tree1 values received to s1:", s1)
	fmt.Println("tree2 values received to s2:", s2)

	fmt.Println("checking that all vals are in both trees")
	for _, elems1 := range s1 {
		for _, elems2 := range s2 {
			if elems1 == elems2 {
				fmt.Println("found match for:", elems1, "in tree2")
				b = true
				break
			}
		}
		if !b {
			fmt.Println("value: ", elems1, " found to be in tree1 but not tree2")
			return false
		}
		b = false
	}
	return true
}

func main() {

	var wg sync.WaitGroup
	var ch1, ch2 chan int = make(chan int), make(chan int, 10)

	wg.Add(1)
	go func() { Walk(tree.New(1), ch1); close(ch1); wg.Done() }()
	for i := range ch1 {
		fmt.Println("received:", i)
	}
	wg.Wait()

	fmt.Println("------------starting Walk(tree.New(10), ch2)----------------")
	wg.Add(1)
	go func() { Walk(tree.New(10), ch2); close(ch2); wg.Done() }()
	for j := range ch2 {
		fmt.Println("received:", j)
	}
	wg.Wait()
	fmt.Println("------------comparing dem trees----------------")
	fmt.Println("these trees are equal?:", Same(tree.New(1), tree.New(1)), "- answer should be true")
	fmt.Println("these trees are equal?:", Same(tree.New(1), tree.New(3)), "- answer should be false")

}
