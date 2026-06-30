package main

import "fmt"

/*
Необходимо написать функцию fruitMarket, которая будет принимать
название фруктов, а возвращать их количество. Например,
«апельсины» - 5. Сама функция должна создать карту: апельсин=5,
яблоки=3, сливы=1, груши=0. Если запрашиваемых фруктов нет в
карте, то в консоль должно выводится сообщение об отсутствии.
*/

func fruitMarket(fruit string) (int, error) {
	var err error
	m := map[string]int{
		"апельсины": 5,
		"яблоки":    3,
		"сливы":     1,
		"груши":     0,
	}

	if _, ok := m[fruit]; !ok {
		err = fmt.Errorf("Фрукта нет")
	}
	return m[fruit], err
}

func main() {
	f, err := fruitMarket("апельсины")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}
