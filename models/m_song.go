package models

// Song represent an object
type Song struct {
	ID       int    `json:"id"`
	UID      string `json:"uid"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Language string `json:"language"`
	Lyric    string `json:"lyric,omitempty"`
}

type LastId struct {
	ID string `json:"id"`
}
