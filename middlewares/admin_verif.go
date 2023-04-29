package middlewares

import (
	"net/http"
	"os"

	"github.com/ardin2001/go_mini-capstone/helpers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func AdminVerification(c echo.Context) (error, bool) {
	err := GetDataJWT(c)
	godotenv.Load()
	role := os.Getenv("ROLE_A")
	if err.Role != role {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "Cannot access this route",
			Status:  false,
		}), false
	}
	return nil, true
}
