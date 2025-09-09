package repository

import (
	"database/sql"
	"notes-app/internal/models"
)

func CreateUser(db *sql.DB, user models.User) error {
	_, err := db.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		user.Username, user.Password,
	)
	return err
}	

func GetUserByUsername(db *sql.DB, username string) (models.User, error) {
	var user models.User
	err := db.QueryRow(
		"SELECT id, username, password FROM users WHERE username=$1",
		username,
	).Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

func GetUserByID(db *sql.DB, id string) (models.User, error) {
	var user models.User
	err := db.QueryRow(
		"SELECT id, username FROM users WHERE id=$1",
		id,
	).Scan(&user.ID, &user.Username)
	return user, err
}

func GetTotalNotesCount(db *sql.DB) (int, error) {
	var total int
	err := db.QueryRow("SELECT COUNT(*) FROM notes").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetTotalNotesByUserIDCount(db *sql.DB,userID string) (int, error) {
	var total int
	err := db.QueryRow("SELECT COUNT(*) FROM notes WHERE user_id = $1", userID).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}