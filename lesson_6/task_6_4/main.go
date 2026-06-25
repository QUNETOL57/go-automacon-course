package main

import "fmt"

/*
Необходимо убрать повторяющийся код - поля Addss и Phone из
структур:
type user struct {
	ID int
	Name string
	Addss string
	Phone string
}
type employee struct {
	ID int
	Name string
	Addss string
	Phone string
}
После проведения рефакторинга строка fmt.Println(u.Addss, u.Phone,
e.Addss, e.Phone) должна выводить в консоль адрес и телефон
пользователя и сотрудника соответственно.
*/

type user struct {
	ID    int
	Name  string
	Addss string
	Phone string
}
type employee struct {
	ID int
	user
}

func main() {
	u := user{
		ID:    1,
		Name:  "Anton",
		Addss: "Russia",
		Phone: "111111",
	}
	e := employee{
		ID: 0,
		user: user{
			ID:    2,
			Name:  "Elena",
			Addss: "Sweden",
			Phone: "222222",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
