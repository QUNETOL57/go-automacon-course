package main

import (
	"lesson_12/internal/entity"
	"lesson_12/internal/service/format"
	"lesson_12/internal/service/report"
)

func main() {
	f := format.Xml{}
	r := report.Report{}
	data := make([]entity.User, 0, 1)
	r.Do(f, data)
}
