package domain

import "time"

type User struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Salt string `json:"salt,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

