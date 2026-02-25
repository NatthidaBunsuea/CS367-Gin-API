package models

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Major string  `json:"major"`
	GPA   float64 `json:"gpa"`
}
