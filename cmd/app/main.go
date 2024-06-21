package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/PyMarcus/api_messanger/internal/dto"
	"github.com/PyMarcus/api_messanger/internal/infra/apachkaf"
	"github.com/PyMarcus/api_messanger/internal/infra/repository"
	"github.com/PyMarcus/api_messanger/internal/infra/web"
	"github.com/PyMarcus/api_messanger/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db settings
	db, err := sql.Open("mysql", "root:your_password@tcp(localhost:3306)/api_kafka_db")
	if err != nil {
		log.Println("Fail while connecting with db", err)
	}
	defer db.Close()

	log.Println("database is ok")

	repo := repository.NewStudentRepositoryMYSQL(db)
	createStudentsUseCase := usecase.NewCreateStudentUseCase(repo)
	listStudentsUseCase := usecase.NewListStudentsUseCase(repo)

	// web server
	studentHandlers := web.NewStudentHandler(createStudentsUseCase, listStudentsUseCase)
	router := chi.NewRouter()
	router.Post("/students", studentHandlers.CreateStudentHandler)
	router.Get("/students", studentHandlers.ListStudentsHandler)

	go http.ListenAndServe(":8000", router)
	log.Println("running server...")
	// kafka
	chanMsg := make(chan *kafka.Message)

	go apachkaf.Consume([]string{"students"}, "localhost:9092", chanMsg)
	log.Println("running kafka...")
	for msg := range chanMsg {
		dto := dto.CreateStudentInputDTO{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			log.Println("Fail to unmarshal consumer json")
			continue
		}
		result, err := createStudentsUseCase.Execute(dto)
		if err != nil {
			log.Println("Fail to recover output from create usecase")
			continue
		}
		log.Println(result)

	}

}
