package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk_helper(t, ch)
	close(ch)
}

func walk_helper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk_helper(t.Left, ch)
	ch <- t.Value
	walk_helper(t.Right, ch)
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 {
			return false
		}

		if v1 != v2 {
			return false
		}

		if !ok1 && !ok2 {
			return true
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		x := <-ch
		fmt.Println(x)
	}

	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(3), tree.New(2)))
}
