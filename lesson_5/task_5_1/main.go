package main

import "fmt"

func main() {
	n := "string"
	var u *string = &n
	fmt.Println(u, *u)
}
