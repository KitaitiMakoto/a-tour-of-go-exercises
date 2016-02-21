package main

import "golang.org/x/tour/wc"
import "strings"

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	wc := make(map[string]int)
	for i := 0; i < len(fields); i++ {
		word := fields[i]
		wc[word]++
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
