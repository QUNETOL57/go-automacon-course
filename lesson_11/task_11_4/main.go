package main

import (
	"errors"
	"fmt"
)

/*
Нужно создать, используя оборачивание, ошибку
«ошибка3:ошибка2:ошибка1». Также нужно создать свою ошибку в виде
структуры myFirstError, которая обязательно должна иметь метод Error()
string. Необходимо убедиться, что в созданной цепочке ошибок не было
ошибки типа myFirstError.
*/

var err = errors.New("ошибка1")

type myFirstError struct {
	message string
	status  string
}

func (e myFirstError) Error() string {
	return e.message
}

func main() {
	err2 := fmt.Errorf("ошибка2:%w", err)
	err3 := fmt.Errorf("ошибка2:%w", err2)

	fmt.Println(errors.As(err3, new(myFirstError)))
}
