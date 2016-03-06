package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		ch <- t.Value
		Walk(t.Left, ch)
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 8)
	ch2 := make(chan int, 8)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			return true
		}
	}
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(Same(t1, t1))
	fmt.Println(Same(t1, t2))
}
