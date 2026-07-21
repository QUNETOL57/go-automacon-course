package main

import (
	"fmt"

	"github.com/QUNETOL57/go-automacon-course/lesson_13/task_13_5/format/v2"
)

/*
В рамках задачи будем работать с картотекой известного врача.
Нужно будет написать модуль с несколькими версиями: v1.0.0, v1.1.0,
v2.0.0, v2.1.0. Модуль должен прочитать файл со следующим
содержимым:

{"name":"Ёжик","age":10,"email":"ezh@mail.ru"}
{"name":"Зайчик","age":2,"email":"zayac@mail.ru"}
{"name":"Лисичка","age":3,"email":"alice@mail.ru"}

Модуль должен содержать функцию Do, которая принимает два
строковых параметра: путь файла откуда прочитать данные, путь
файла, в который записать данные в требуемом формате; и
возвращать ошибку. Пример использования модуля:
package main
import (
	format ...
)
func main() {
	err := format.Do("patients", "result")
	if err != nil {
	…
	}
}
Должна быть возможность подключить любую из версий: v3.0.0, v3.1.0,
v4.0.0, v4.1.0.
*/

func main() {
	err := format.Do("inputFile", "outputFile")
	if err != nil {
		fmt.Println(err)
	}
}
