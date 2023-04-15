package models

// PostGet represent
type PostGet struct {
	ID       int    `json:"id"`
	UID      string `json:"uid"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Content  string `json:"content,omitempty"`
	Language string `json:"language"`
	Username string `json:"username,omitempty"`
	Image    string `json:"image,omitempty"`
}

type Post struct {
	ID       int    `json:"id"`
	UID      string `json:"uid"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Content  string `json:"content,omitempty"`
	Language string `json:"language"`
	Hashtag  string `json:"hashtag,omitempty"`
}
