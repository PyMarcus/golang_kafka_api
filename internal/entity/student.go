package entity

import (
	"github.com/google/uuid"
)

type Student struct {
	Id   string `db:"id"`
	Name string `db:"name"`
	Age  uint16 `db:"age"`
}

func NewStudent(name string, age uint16) *Student {
	return &Student{
		Name: name,
		Id:   uuid.New().String(),
		Age:  age,
	}
}
