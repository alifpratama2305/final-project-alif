package models

// Status represents the model for an status
type Status struct {
	StatusID  int8   `json:"status_id" example:"1"`
	StatusTxt string `json:"status_txt" example:"New"`
}
