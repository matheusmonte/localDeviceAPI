package models

type Device struct {
	ID uint `json:"id" form:"primary_key"`
	Nickname string `json:"nickname"`
	IP string `json:"ip"`
}
