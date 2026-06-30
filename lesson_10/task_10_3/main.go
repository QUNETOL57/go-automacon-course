package main

import (
	f "first"
	s "second"

	"fmt"
)

/*
Необходимо создать и установить пакет second. Пакет должен иметь
функцию Hello, которая возвращает строку «Hello, Louis!». Далее
написать программу, которая будет вызывать функцию Hello пакета first
задачи 10_1 и функцию Hello пакета second. Оба результата должны
выводиться в консоль.
*/

func main() {
	fmt.Println(f.Hello())
	fmt.Println(s.Hello())
}
