package models

// Comment represent comment
type Comment struct {
	ID        int    `json:"id"`
	PostId    string `json:"post_id"`
	UID       string `json:"uid"`
	Comment   string `json:"comment"`
	Timestamp string `json:"timestamp"`
}
