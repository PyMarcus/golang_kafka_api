package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PyMarcus/api_messanger/internal/dto"
	"github.com/PyMarcus/api_messanger/internal/usecase"
)

type StudentHandler struct {
	CreateStudentUseCase *usecase.CreateStudentUseCase
	ListStudentsUseCase  *usecase.ListStudentsUseCase
}

func NewStudentHandler(createUseCase *usecase.CreateStudentUseCase, listUseCase *usecase.ListStudentsUseCase) *StudentHandler {
	return &StudentHandler{createUseCase, listUseCase}
}

func (s StudentHandler) CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var body dto.CreateStudentInputDTO

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("Fail while decode body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, _ := s.CreateStudentUseCase.Execute(body)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s StudentHandler) ListStudentsHandler(w http.ResponseWriter, r *http.Request) {
	response, err := s.ListStudentsUseCase.Execute()

	if err != nil {
		log.Println("Internal server error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
