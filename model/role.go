package model

import "github.com/google/uuid"

type Role struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"login"`
}

func (r *Role) CreateId() {
	r.ID = uuid.New()
}
