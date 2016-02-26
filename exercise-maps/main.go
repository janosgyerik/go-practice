package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)

	for _, word := range strings.Fields(s) {
		counts[word]++
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
