package main

import (
	"encoding/xml"
	"fmt"
)

/*
Необходимо представить в виде xml структуру contracts
type contracts struct {
	List []contract `xml:"contract"`
}
contract{
	Number: 1,
	Landlord: "Остап Бендер",
	tenat: "Шура Балаганов",
}
*/

type contracts struct {
	List []contract `xml:"contract"`
}

type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	tenat    string `xml:"tenat"`
}

func main() {
	c := contract{
		Number:   1,
		Landlord: "Остап Бендер",
		tenat:    "Шура Балаганов",
	}

	r, err := xml.Marshal(contracts{List: []contract{c}})
	if err != nil {
		return
	}
	fmt.Printf("%s", r)
}
