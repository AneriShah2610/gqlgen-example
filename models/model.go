package models

// Job Model
type Job struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	CreatedBy   string `json:"createdBy"`
}

// NewJob to create New Job
type NewJob struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	CreatedByID string `json:"createdByID"`
}
