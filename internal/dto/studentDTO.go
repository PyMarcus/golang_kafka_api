package dto

type CreateStudentInputDTO struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

type CreateStudentOutPutDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}
