package format

import (
	"lesson_12/internal/entity"
	"log"
)

type Xml struct {
	version string
	root    string
}

func (x Xml) Show(d []entity.User) (string, error) {
	log.Println("данные в формате xml")
	return "данные в формате xml", nil
}
