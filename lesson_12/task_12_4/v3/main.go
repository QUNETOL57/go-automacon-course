package main

import (
	"errors"
	"fmt"
)

/*
В представленном ниже коде возникает паника. Нужно понять где и
объяснить почему. Также нужно предложить не менее 3-х вариантов
решения проблемы с объяснением причины устранения ошибки.
*/

type Bird interface {
	Sing() string
}
type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	return d.voice
}

func main() {
	var d *Duck = new(Duck)
	d.voice = "krua"
	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}
