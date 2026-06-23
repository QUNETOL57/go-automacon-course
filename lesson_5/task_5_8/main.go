package main

import "fmt"

type square int

func main() {
	s := square(34)
	s += 10
	fmt.Printf("%d м²", s)
}
