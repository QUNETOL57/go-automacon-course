package main

import "fmt"

func main() {
	first := 16
	second := 3

	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %T.\n", first/second, first%second, first/second)
}
