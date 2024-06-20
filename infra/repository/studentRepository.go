package repository

import "github.com/PyMarcus/api_messanger/entity"

type StudentRepository interface {
	Create(*entity.Student) error
	FindAll() ([]*entity.Student, error)
}
