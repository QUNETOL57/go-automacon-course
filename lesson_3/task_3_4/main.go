package main

import "fmt"

const (
	myConst0 = iota
	myConst1
	myConst2
	myConst3
	myConst4
)

func main() {
	fmt.Println(myConst0, myConst1, myConst2, myConst3, myConst4)
}
