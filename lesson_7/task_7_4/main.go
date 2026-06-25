package main

import "fmt"

func main() {
	s := make([]int, 0, 6)
	s = append(s, 5, 2, 8, 3, 1, 9)
	fmt.Println(s)
}
