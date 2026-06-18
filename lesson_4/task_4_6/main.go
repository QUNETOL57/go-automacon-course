package main

import "fmt"

const count = 23

func fib(n int, m int, i int) {
	var f int
	if i > count {
		return
	}

	if i == 0 || i == 1 {
		f = i
	} else {
		f = n + m
	}

	fmt.Println(f)
	fib(m, f, i+1)
}

func main() {
	fib(0, 0, 0)
}
