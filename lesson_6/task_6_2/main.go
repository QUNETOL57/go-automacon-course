package main

import "fmt"

/*
Необходимо объявить локальную структуру contract с полями: ID int,
Number string, Date string. Далее создать экземпляр структуры со
значениями полей: ID=1, Number=«#000A101\t01», Date=«2024-01-31». В
консоль нужно вывести структуру таким образом, чтобы данные
отображались в виде: {ID:1 Number:#000A101 01 Date:2024-01-31}
*/

func main() {
	type contract struct {
		ID     int
		Number string
		Date   string
	}
	c := contract{1, "#000A101\t01", "2024-01-31"}

	fmt.Printf("%+v", c)
}
