package models

// Artist represent an object
type Artist struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
