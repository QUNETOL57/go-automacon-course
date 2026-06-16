package main

import "fmt"

func main() {
	var first int = 16
	var second int = 3

	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %T.\n", first/second, first%second, first/second)
}
