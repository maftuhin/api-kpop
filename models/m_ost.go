package models

// Track represent an object
type Track struct {
	ID    string `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

// Link represent an object
type Link struct {
	ID         int    `json:"id"`
	Soundtrack string `json:"soundtrack"`
	Song       string `json:"song"`
	Artist     string `json:"artist"`
}
