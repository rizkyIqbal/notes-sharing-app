package repository

import (
	"database/sql"
	"time"
)

func CreateRefreshToken(db *sql.DB, userID string, token string, expiresAt time.Time) error {
	_, err := db.Exec(`
		INSERT INTO refresh_tokens (user_id, token, expires_at)
		VALUES ($1, $2, $3)
	`, userID, token, expiresAt)
	return err
}

func GetRefreshToken(db *sql.DB, token string) (string, string, time.Time, error) {
	var id string
	var userID string
	var expiresAt time.Time
	err := db.QueryRow(`
		SELECT id, user_id, expires_at
		FROM refresh_tokens
		WHERE token = $1 AND revoked_at IS NULL
	`, token).Scan(&id, &userID, &expiresAt)
	return id, userID, expiresAt, err
}

func RevokeRefreshToken(db *sql.DB, token string) error {
	_, err := db.Exec(`
		UPDATE refresh_tokens
		SET revoked_at = NOW()
		WHERE token = $1
	`, token)
	return err
}

func RevokeUserTokens(db *sql.DB, userID string) error {
	_, err := db.Exec(`
		UPDATE refresh_tokens
		SET revoked_at = NOW()
		WHERE user_id = $1
	`, userID)
	return err
}