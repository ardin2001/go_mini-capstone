package routes

import (
	"net/http"

	"github.com/ardin2001/go_mini-capstone/configs"
	"github.com/ardin2001/go_mini-capstone/controllers"
	"github.com/ardin2001/go_mini-capstone/repositories"
	"github.com/ardin2001/go_mini-capstone/services"
	"github.com/labstack/echo/v4"
)

var (
	db, _ = configs.InitDB()
	userR = repositories.NewUserRepositories(db)
	userS = services.NewUserServices(userR)
	userC = controllers.NewUserControllers(userS)
)

func StartApp() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", userC.GetUserController)
	e.GET("/users", userC.GetUsersController)
	e.GET("users/auth", userC.GetUsersController)
	e.POST("/users", userC.CreateUserController)
	e.DELETE("/users/:id", userC.DeleteUserController)
	e.PUT("/users/:id", userC.UpdateUserController)

	return e
}
