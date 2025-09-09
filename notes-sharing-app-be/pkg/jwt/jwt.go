package jwt

import (
	"errors"
	"time"

	"notes-app/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string) (string, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(config.JWTPrivateKey)
	if err != nil {
		return "", err
	}

	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
	}

	// claims := jwt.MapClaims{
    //     "user_id": userID,
    //     "exp":     time.Now().Add(45 * time.Minute).Unix(),
    // }

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

func ValidateToken(tokenStr string) (string, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(config.JWTPublicKey)
	if err != nil {
		return "", err
	}

	// token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	// claims, ok := token.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return "", errors.New("invalid claims")
	// }

	// userID, ok := claims["user_id"].(float64)
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", errors.New("user_id not found in token")
	}

	return claims.Subject, nil
}
