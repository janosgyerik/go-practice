package main

import "fmt"

func fibonacci() func() int {
	prev, current := 0, 1

	return func() int {
		prev, current = current, prev + current
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
