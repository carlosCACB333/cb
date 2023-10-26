package libs

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(secret string, userID string, exp time.Duration) (string, time.Time, error) {
	expirationTime := time.Now().Add(exp)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secret))
	if err != nil {
		return "", expirationTime, err
	}
	return token, expirationTime, nil
}

func ValidateToken(secret string, token string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	userID := claims["user_id"].(string)
	return userID, nil
}
