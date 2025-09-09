package service

import (
	"database/sql"
	"errors"
	"notes-app/internal/repository"
	"notes-app/pkg/jwt"
	"time"

	"github.com/google/uuid"
)

func GenerateRefreshToken(db *sql.DB, userID string) (string, error) {
	token := uuid.NewString() 
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	if err := repository.CreateRefreshToken(db, userID, token, expiresAt); err != nil {
		return "", err
	}

	return token, nil
}

func RefreshAccessToken(db *sql.DB, refreshToken string) (string, error) {
	_, userID, expiresAt, err := repository.GetRefreshToken(db, refreshToken)
	if err != nil {
		return "", err
	}

	if time.Now().After(expiresAt) {
		return "", errors.New("refresh token expired")
	}

	// generate new access token
	accessToken, err := jwt.GenerateToken(userID)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func RevokeRefreshTokenService(db *sql.DB, token string) error {
	return repository.RevokeRefreshToken(db, token)
}
