package main

import "fmt"

func main() {
	kart := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	m := [3]string{"слон", "бегемот", "осьминог"}

	for _, v := range m {
		count, inMap := kart[v]
		fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", v, count, inMap)
	}
}
