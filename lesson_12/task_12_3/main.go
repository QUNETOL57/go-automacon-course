package main

import "fmt"

/*
Не изменяя структуры Xml, Csv и функцию main необходимо
доработать следующий код так, чтобы в консоли увидели:
Данные в формате xml
Данные в формате csv
*/

type File interface {
	Format()
}

type Xml struct{}

func (x Xml) Format() {
	fmt.Println("Данные в формате xml")
}

type Csv struct{}

func (c Csv) Format() {
	fmt.Println("Данные в формате csv")
}

func main() {
	x := Xml{}
	Report(x)
	c := Csv{}
	Report(c)
}

func Report(x File) {
	x.Format()
}
