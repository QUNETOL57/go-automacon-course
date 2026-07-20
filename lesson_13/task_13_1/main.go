package main

import (
	"encoding/json"
	"fmt"
)

/*
Необходимо распарсить json {"number":1,"landlord":"Остап
Бендер","tenant":"Шура Балаганов»}
*/

type data struct {
	Number   int
	Landlord string
	Tenant   string
}

func main() {
	s := `{"number":1,"landlord":"Остап\nБендер","tenant":"Шура Балаганов"}`
	v := data{}
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		return
	}
	fmt.Printf("%+v", v)
}
