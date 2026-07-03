package main

import (
	"fmt"
	"io"
	"os"
)

type myStruct struct{}

func (m myStruct) Write(b []byte) (n int, err error) {
	return 0, nil
}
func main() {
	w1 := new(myStruct)
	runW(w1)
	w2 := os.Stdout
	runW(w2)
	w3 := myStruct{}
	runW(w3)
}
func runW(w io.Writer) {
	switch w.(type) {
	case myStruct:
		fmt.Println("это myStruct")
	case *os.File:
		fmt.Println("это *os.File")
	default:
		fmt.Println("неизвестный тип")
	}
}
