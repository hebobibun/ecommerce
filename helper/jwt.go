package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateToken(role string, userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role":       role,
		"id":         userID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
		"authorized": true,
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
