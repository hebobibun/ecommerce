package model

import "github.com/google/uuid"

type Admin struct {
	ID       uuid.UUID `gorm:"type:uuid; primaryKey"`
	Username string    `json:"username" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required"`
	Token    string    `json:"token"`
	Model
}
