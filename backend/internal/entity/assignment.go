package entity

type (
	Assignment struct {
		ID        string `json:"id"`
		Subject   string `json:"subject"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		StudentId string `json:"student_id"`
	}
)
