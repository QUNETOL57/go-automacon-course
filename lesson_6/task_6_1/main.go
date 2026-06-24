package main

import "fmt"

type contact struct {
	ID     int
	Number string
	Date   string
}

func main() {
	c := contact{1, "#000A\n101", "2024-01-31"}
	fmt.Printf("%#v", c)
}
