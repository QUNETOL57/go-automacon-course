package main

import (
	"errors"
	"fmt"
)

/*
Нужно создать, используя оборачивание, ошибку
«ошибка3:ошибка2:ошибка1». Из созданной цепочки ошибок нужно
получить ошибку «ошибка2:ошибка1» и вывести в stdout.
*/

func main() {
	err := errors.New("ошибка1")

	err = fmt.Errorf("ошибка2:%w", err)
	err = fmt.Errorf("ошибка3:%w", err)
	err = errors.Unwrap(err)
	fmt.Println(err)
}
