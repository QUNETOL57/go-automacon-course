package main

import "fmt"

func main() {
	m := [4]string{"яблоко", "груша", "помидор", "абрикос"}
	m[2] = "персик"
	fmt.Println(m)
}
