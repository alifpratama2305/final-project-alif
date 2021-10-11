package models

// Status represents the model for an status
type Status struct {
	StatusID  int8   `json:"statusId" example:"1"`
	StatusTxt string `json:"statusTxt" example:"New"`
}
