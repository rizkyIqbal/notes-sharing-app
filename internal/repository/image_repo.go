package repository

import (
	"database/sql"
	"fmt"
	"notes-app/internal/models"
)

func CreateImages(db *sql.DB, noteID string, urls []string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	stmt, err := tx.Prepare(`INSERT INTO images (note_id, image_path) VALUES ($1, $2)`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, url := range urls {
		_, err := stmt.Exec(noteID, url)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert image (%s): %w", url, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func GetImagesByNoteID(db *sql.DB, noteID string) ([]models.Image, error) {
	rows, err := db.Query(
		`SELECT id, note_id, image_path, created_at, updated_at 
		 FROM images 
		 WHERE note_id = $1`, noteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []models.Image
	for rows.Next() {
		var img models.Image
		if err := rows.Scan(&img.ID, &img.NoteID, &img.ImagePath, &img.CreatedAt, &img.UpdatedAt); err != nil {
			return nil, err
		}
		images = append(images, img)
	}
	return images, nil
}

func DeleteImage(db *sql.DB, imageID string) error {
    query := `DELETE FROM images WHERE id = $1`
    _, err := db.Exec(query, imageID)
    return err
}