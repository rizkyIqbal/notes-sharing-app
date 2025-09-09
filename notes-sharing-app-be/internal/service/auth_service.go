package service

import (
	"database/sql"
	"errors"
	"notes-app/internal/models"
	"notes-app/internal/repository"
	"notes-app/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sql.DB, user models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return repository.CreateUser(db, user)
}

func LoginUser(db *sql.DB, username, password string) (string, error) {
	user, err := repository.GetUserByUsername(db, username)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUserProfile(db *sql.DB, userID string) (models.User, error) {
	return repository.GetUserByID(db, userID)
}