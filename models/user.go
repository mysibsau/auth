package models

type User struct {
	Id       string `json:"id"`
	Group    Group  `json:"group"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	password string
}
