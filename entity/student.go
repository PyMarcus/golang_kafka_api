package entity

import (
	"github.com/google/uuid"
)

type Student struct {
	Name string
	Id   string
	Age  uint16
}

func NewStudent(name string, age uint16) *Student {
	return &Student{
		Name: name,
		Id:   uuid.New().String(),
		Age:  age,
	}
}
