package models

import "time"

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" validate:"required,min=5"`
	Content   string    `json:"content" validate:"required,min=5"`
	UserID	  string	`json:"user_id" validate:"required"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
