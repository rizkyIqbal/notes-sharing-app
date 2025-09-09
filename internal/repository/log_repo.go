package repository

import (
	"database/sql"
	"notes-app/internal/models"
)

func InsertLog(db *sql.DB, log models.Log) error {
	_, err := db.Exec(
		`INSERT INTO logs (datetime, method, endpoint, headers, payload, response_body, status_code) 
         VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		log.Datetime, log.Method, log.Endpoint, log.Headers,
		log.Payload, log.ResponseBody, log.StatusCode,
	)
	return err
}
