package repository

import (
	"database/sql"

	"github.com/PyMarcus/api_messanger/entity"
)

type StudentRepositoryMYSQL struct {
	DB *sql.DB
}

func NewStudentRepositoryMYSQL(db *sql.DB) StudentRepository {
	return StudentRepositoryMYSQL{db}
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
	rows, err := s.DB.Query("SELECT * FROM students;")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student *entity.Student
		rows.Scan(student)
		students = append(students, student)
	}
	return students, nil
}
