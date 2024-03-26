package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}

func (usr *User) CreateId() {
	usr.ID = uuid.New()
}

type Token struct {
	Value string `json:"token"`
}