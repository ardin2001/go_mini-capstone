package middlewares

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func AdminVerification(c echo.Context) (error, bool) {
	err := GetDataJWT(c)
	godotenv.Load()
	role := os.Getenv("ROLE_A")
	if err.Role != role {
		return errors.New("you do not have rights to access this route"), false
	}
	return nil, true
}
