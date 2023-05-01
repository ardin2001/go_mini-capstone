package middlewares

import (
	"os"
	"time"

	"github.com/ardin2001/go_mini-capstone/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func CreateToken(userId uint, nama string, role string) (string, error) {
	claims := &models.JwtCustomClaims{
		ID:   userId,
		Nama: nama,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 100)),
		},
	}

	// Create token with claims
	godotenv.Load()
	dbHost := os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(dbHost))

}
