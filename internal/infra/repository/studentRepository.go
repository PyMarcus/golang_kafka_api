package repository

import "github.com/PyMarcus/api_messanger/internal/entity"

type StudentRepository interface {
	Create(*entity.Student) error
	FindAll() ([]*entity.Student, error)
}
