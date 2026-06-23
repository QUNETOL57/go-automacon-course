package main

import "fmt"

func change(n *int) {
	*n = *n + 1
}

func main() {
	n := 10
	change(&n)
	fmt.Println(n)
}
