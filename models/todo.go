package models

// Todo represents the model for an todo
type Todo struct {
	Title          string `json:"title" example:"title"`
	Description    string `json:"description" example:"description"`
	DueDate        string `json:"dueDate" example:"2021-10-02"`
	PersonInCharge string `json:"personInCharge" example:"alif"`
	Status         string `json:"status" example:"New"`
}
