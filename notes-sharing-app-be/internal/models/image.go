package models

import "time"

type Image struct {
	ID        string   	`json:"id"`
	NoteID   	string    	`json:"note_id" validate:"required,min=5"`
	ImagePath   string  `json:"image_path" validate:"required,min=5"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

