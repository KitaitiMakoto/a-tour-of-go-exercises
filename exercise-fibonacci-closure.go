package main

import "fmt"

func fibonacci() func() int {
	i := 0
	j := 1
	f := func() int {
		i, j = j, i + j
		return i
	}
	return f
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
