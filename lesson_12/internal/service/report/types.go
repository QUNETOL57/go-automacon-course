package report

import "lesson_12/internal/entity"

type Formatter interface {
	Show(d []entity.User) (string, error)
}
