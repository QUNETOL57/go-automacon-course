package main

import "fmt"

func hello(f func()) {
	f()
}

func main() {
	hello(func() {
		fmt.Println("Hello, Go!")
	})
}
