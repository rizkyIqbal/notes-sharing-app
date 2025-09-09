package repository

import (
	"database/sql"
	"notes-app/internal/models"
)

func CreateNote(db *sql.DB, note models.Note) (string, error) {
	var id string
	err := db.QueryRow(
		"INSERT INTO notes (title, content, user_id) VALUES ($1, $2, $3) RETURNING id",
		note.Title, note.Content, note.UserID,
	).Scan(&id)
	return id, err
}

func GetNoteByID(db *sql.DB, id string) (models.Note, error) {
	var note models.Note
	err := db.QueryRow(
		"SELECT n.id, n.title, n.content, n.user_id, u.username, n.created_at, n.updated_at FROM notes n JOIN users u ON n.user_id = u.id WHERE n.id = $1",
		id,
	).Scan(&note.ID,
        &note.Title,
        &note.Content,
        &note.UserID,
        &note.Username,
        &note.CreatedAt,
        &note.UpdatedAt,)
	return note, err
}

func GetNotesByUserID(db *sql.DB, userID, title string, limit, offset int) ([]models.Note, error) {
	var rows *sql.Rows
	var err error

	if title != "" {
		query := `
			SELECT id, title, content, user_id, created_at, updated_at
			FROM notes
			WHERE user_id = $1 AND title ILIKE $2
			ORDER BY created_at DESC
			LIMIT $3 OFFSET $4
		`
		rows, err = db.Query(query, userID, "%"+title+"%", limit, offset)
	} else {
		query := `
			SELECT id, title, content, user_id, created_at, updated_at
			FROM notes
			WHERE user_id = $1
			ORDER BY created_at DESC
			LIMIT $2 OFFSET $3
		`
		rows, err = db.Query(query, userID, limit, offset)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var n models.Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.UserID, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}

	return notes, nil
}



func GetAllNotes(db *sql.DB, title string, limit, offset int) ([]models.Note, error) {
    var rows *sql.Rows
    var err error

    if title != "" {
        query := `
            SELECT id, title, content, user_id, created_at, updated_at
            FROM notes
            WHERE title ILIKE $1
            ORDER BY created_at DESC
            LIMIT $2 OFFSET $3
        `
        rows, err = db.Query(query, "%" + title + "%", limit, offset)
    } else {
        query := `
            SELECT id, title, content, user_id, created_at, updated_at
            FROM notes
            ORDER BY created_at DESC
            LIMIT $1 OFFSET $2
        `
        rows, err = db.Query(query, limit, offset)
    }

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var notes []models.Note
    for rows.Next() {
        var n models.Note
        if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.UserID, &n.CreatedAt, &n.UpdatedAt); err != nil {
            return nil, err
        }
        notes = append(notes, n)
    }

    return notes, nil
}


func UpdateNote(db *sql.DB, note models.Note) error {
	_, err := db.Exec(
		`UPDATE notes 
		 SET title = $1, content = $2, updated_at = NOW() 
		 WHERE id = $3 AND user_id = $4`,
		note.Title, note.Content, note.ID, note.UserID,
	)
	return err
}

func DeleteNote(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM notes WHERE id=$1", id)
	return err
}
