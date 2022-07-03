package models

type User struct {
	Id       string  `json:"id"`
	Group    *string `json:"group"`
	Name     *string `json:"name"`
	Login    string  `json:"login"`
	password string
	Average  float64 `json:"average"`
}
