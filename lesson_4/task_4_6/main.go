package main

import "fmt"

const (
	n = iota
	m
	count = 23
)

func fibonacci(n int, m int, count int) {
	if count <= 0 {
		return
	}
	f := n + m
	fmt.Printf(" %d", f)
	fibonacci(m, f, count-1)
}

func main() {
	fmt.Print(n, m)
	fibonacci(n, m, count)
}
