package models

type User struct {
	ID          int    `json:"id"`
	UID         string `json:"uid"`
	Image       string `json:"image"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Email       string `json:"email"`
}
