package models

// Todo represents the model for an todo
type Todo struct {
	Title          string `json:"title" example:"title"`
	Description    string `json:"description" example:"description"`
	DueDate        string `json:"due_date" example:"2021-10-02"`
	PersonInCharge string `json:"person_in_charge" example:"dwi"`
	Status         string `json:"status" example:"New"`
}
