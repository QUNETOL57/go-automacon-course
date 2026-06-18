package main

import "fmt"

func hello(lang string) (str string) {
	str = fmt.Sprintf("Hello, %s!", lang)
	return
}

func main() {
	fmt.Println(hello("Go"))
}
