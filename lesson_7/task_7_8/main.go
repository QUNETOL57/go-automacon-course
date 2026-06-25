package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	s = append(s[:4], s[5:]...)

	fmt.Println(s)
}
