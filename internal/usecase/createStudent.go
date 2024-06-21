package usecase

import (
	"github.com/PyMarcus/api_messanger/internal/dto"
	"github.com/PyMarcus/api_messanger/internal/entity"
	"github.com/PyMarcus/api_messanger/internal/infra/repository"
)

type CreateStudentUseCase struct {
	StudentRepository repository.StudentRepository
}

func (c CreateStudentUseCase) Execute(input dto.CreateStudentInputDTO) (dto.CreateStudentOutPutDTO, error) {
	student := entity.NewStudent(input.Name, input.Age)
	err := c.StudentRepository.Create(student)
	if err != nil {
		return dto.CreateStudentOutPutDTO{}, err
	}
	return dto.CreateStudentOutPutDTO{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func NewCreateStudentUseCase(sr repository.StudentRepository) *CreateStudentUseCase {
	return &CreateStudentUseCase{sr}
}
