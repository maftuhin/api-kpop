package models

// Track represent an object
type Track struct {
	ID    int    `json:"id"`
	UID   string `json:"uid"`
	Title string `json:"title"`
	Image string `json:"image"`
}
