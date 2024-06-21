package repository

import (
	"database/sql"
	"log"

	"github.com/PyMarcus/api_messanger/internal/entity"
)

type StudentRepositoryMYSQL struct {
	DB *sql.DB
}

func NewStudentRepositoryMYSQL(db *sql.DB) *StudentRepositoryMYSQL {
	return &StudentRepositoryMYSQL{db}
}

func (s StudentRepositoryMYSQL) Create(student *entity.Student) error {
	_, err := s.DB.Exec("INSERT INTO students (id, name, age) VALUES (?, ?, ?)", student.Id, student.Name, student.Age)
	if err != nil {
		return err
	}
	return nil
}

func (s StudentRepositoryMYSQL) FindAll() ([]*entity.Student, error) {
	var students []*entity.Student

	rows, err := s.DB.Query("SELECT id, name, age FROM students;")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student entity.Student
		if err := rows.Scan(&student.Id, &student.Name, &student.Age); err != nil {
			return nil, err
		}
		log.Print("nome", student.Name)
		students = append(students, &student)
	}
	return students, nil
}
