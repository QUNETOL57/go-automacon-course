package main

import (
	"errors"
	"fmt"
)

/*
Нужно создать, используя оборачивание, ошибку
«ошибка3:ошибка2:ошибка1». Не используя unwrap, нужно определить
была ли ошибка «ошибка1»
*/
var err = errors.New("ошибка1")

func main() {
	err2 := fmt.Errorf("ошибка2:%w", err)
	err3 := fmt.Errorf("ошибка3:%w", err2)

	fmt.Println(errors.Is(err3, err))
}
