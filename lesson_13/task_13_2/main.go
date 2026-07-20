package main

import (
	"encoding/json"
	"fmt"
)

/*
Необходимо представить в виде json структуру contract
contract{
	Number: 2,
	Landlord: "Остап",
	tenat: "Шура",
}
*/

type contract struct {
	Number   int
	Landlord string
	tenant   string
}

func main() {
	c := contract{
		Number:   2,
		Landlord: "Остап",
		tenant:   "Шура",
	}
	r, err := json.Marshal(c)
	if err != nil {
		return
	}
	fmt.Println(string(r))
}
