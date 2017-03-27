package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func walkTreeRecurse(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkTreeRecurse(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkTreeRecurse(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkTreeRecurse(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2, ch2Ok, j := make(chan int), make(chan int), true, 0
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if ch2Ok {
			j, ch2Ok = <-ch2
			if i != j {
				return false
			}
		} else {
			return false
		}
	}
	_, ch2Ok = <-ch2
	if ch2Ok {
		return false
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(100), tree.New(100)))
}
