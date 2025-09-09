package service

import (
	"database/sql"
	"notes-app/internal/models"
	"notes-app/internal/repository"
)

func AddImagesToNote(db *sql.DB, noteID string, urls []string) error {
	return repository.CreateImages(db, noteID, urls)
}

func GetImages(db *sql.DB, noteID string) ([]models.Image, error) {
	return repository.GetImagesByNoteID(db, noteID)
}

func DeleteImageService(db *sql.DB, imageID string) error {
    return repository.DeleteImage(db, imageID)
}
