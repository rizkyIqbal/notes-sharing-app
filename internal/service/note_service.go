package service

import (
	"database/sql"
	"notes-app/internal/models"
	"notes-app/internal/repository"
)

func GetAllNotesService(db *sql.DB, title string, limit int, page int) ([]models.Note, error) {
	return repository.GetAllNotes(db, title, limit, page)
}

func GetNotesByUserIDService(db *sql.DB, id string, title string, limit int, page int) ([]models.Note, error) {
	return repository.GetNotesByUserID(db, id, title, limit, page)
}

func CreateNoteService(db *sql.DB, note models.Note) (string, error) {
	return repository.CreateNote(db, note)
}

func GetNoteService(db *sql.DB, id string) (models.Note, error) {
	return repository.GetNoteByID(db, id)
}

func GetTotalNotesCountService(db *sql.DB) (int, error){
	return repository.GetTotalNotesCount(db)
}

func GetTotalNotesByUserIDCountService(db *sql.DB, userID string) (int, error){
	return repository.GetTotalNotesByUserIDCount(db, userID)
}

func UpdateNoteService(db *sql.DB, note models.Note) error {
	return repository.UpdateNote(db, note)
}

func DeleteNoteService(db *sql.DB, id string) error {
	return repository.DeleteNote(db, id)
}
