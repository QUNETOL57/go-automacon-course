package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

/*
Необходимо распарсить xml
<?xml version="1.0" encoding="UTF-8"?>
<contracts>
    <contract>
        <number>1</number>
        <sign_date>2023-09-02</sign_date>
        <landlord>Остап</landlord>
        <tenat>Шура</tenat>
    </contract>
    <contract>
        <number>2</number>
        <sign_date>2023-09-03</sign_date>
        <landlord>Бендер</landlord>
        <tenat>Балаганов</tenat>
    </contract>
</contracts>
*/

type contract struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

type contracts struct {
	List []contract `xml:"contract"`
}

func main() {
	f, err := os.Open("lesson_13/task_13_3/1.xml")
	if err != nil {
		return
	}
	defer f.Close()
	r := contracts{}
	err = xml.NewDecoder(f).Decode(&r)
	if err != nil {
		return
	}
	fmt.Printf("%+v", r)
}
