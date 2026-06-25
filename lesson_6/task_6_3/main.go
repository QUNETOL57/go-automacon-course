package main

import "fmt"

/*
Необходимо объявить глобальную структуру contract с полями: ID int,
Number string, Date string. Далее создать экземпляр структуры со
значениями полей: ID=1, Number=«#000A\n101», Date=«2024-01-31». При
передачи экземпляра структуры в fmt.Println в консоли должно
отображаться: Договор № #000A\n101 от 2024-01-31
*/

type contact struct {
	ID           int
	Number, Date string
}

func (c contact) print() string {
	return fmt.Sprintf("Договор № %#v от %s", c.Number, c.Date)
}

func main() {
	c := contact{1, "#000A\n101", "2024-01-31"}

	fmt.Println(c.print())
}
