package models

type Session struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Mail     string `json:"mail"`
}
