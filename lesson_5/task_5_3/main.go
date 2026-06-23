package main

import "fmt"

func main() {
	b := false
	var u *bool = &b

	*u = true
	fmt.Println(b)
}
