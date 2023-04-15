package models

type Language struct {
	ID       int    `json:"id"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

type Request struct {
	UID    string `json:"uid"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}
