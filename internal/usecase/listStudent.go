package usecase

import (
	"github.com/PyMarcus/api_messanger/internal/dto"
	"github.com/PyMarcus/api_messanger/internal/infra/repository"
)

type ListStudentsUseCase struct {
	StudentRepository repository.StudentRepository
}

func (l ListStudentsUseCase) Execute() ([]*dto.ListStudentsOutPutDTO, error) {
	students, err := l.StudentRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var studentsOutputDTO []*dto.ListStudentsOutPutDTO
	for _, s := range students {
		studentsOutputDTO = append(studentsOutputDTO, &dto.ListStudentsOutPutDTO{
			Id:   s.Id,
			Name: s.Name,
			Age:  s.Age,
		})
	}

	return studentsOutputDTO, nil
}

func NewListStudentsUseCase(sr repository.StudentRepository) *ListStudentsUseCase {
	return &ListStudentsUseCase{
		sr,
	}
}
