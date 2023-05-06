package routes

import (
	"net/http"
	"os"

	"github.com/ardin2001/go_mini-capstone/configs"
	"github.com/ardin2001/go_mini-capstone/controllers"
	"github.com/ardin2001/go_mini-capstone/middlewares"
	"github.com/ardin2001/go_mini-capstone/models"
	"github.com/ardin2001/go_mini-capstone/repositories"
	"github.com/ardin2001/go_mini-capstone/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	db, _ = configs.InitDB()
	userR = repositories.NewUserRepositories(db)
	userS = services.NewUserServices(userR)
	userC = controllers.NewUserControllers(userS)

	productR = repositories.NewProductRepositories(db)
	productS = services.NewProductServices(productR)
	productC = controllers.NewProductControllers(productS)

	cartR = repositories.NewCartRepositories(db)
	cartS = services.NewCartServices(cartR)
	cartC = controllers.NewCartControllers(cartS)

	transactionDetailR = repositories.NewTransactionDetailRepositories(db)
	transactionDetailS = services.NewTransactionDetailServices(transactionDetailR)
	transactionDetailC = controllers.NewTransactionDetailControllers(transactionDetailS)

	transactionR = repositories.NewTransactionRepositories(db)
	transactionS = services.NewTransactionServices(transactionR)
	transactionC = controllers.NewTransactionControllers(transactionS, transactionDetailS, cartS)
)

func StartApp() *echo.Echo {
	e := echo.New()
	middlewares.Logger(e)
	godotenv.Load()
	dbHost := os.Getenv("SECRET_KEY")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte(dbHost),
	}
	e.Static("/images", "images/")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// root level, user tidak diperbolehkan akses semua data user
	e.GET("/users", userC.GetUsersController, echojwt.WithConfig(config))
	e.POST("/products", productC.CreateProductController, echojwt.WithConfig(config))
	e.DELETE("/products/:id", productC.DeleteProductController, echojwt.WithConfig(config))
	e.PUT("/products/:id", productC.UpdateProductController, echojwt.WithConfig(config))
	e.GET("/carts", cartC.GetCartsController, echojwt.WithConfig(config))
	// admin

	e.POST("/users/login", userC.LoginUserController)
	e.GET("/users/me", userC.GetUserController, echojwt.WithConfig(config))
	e.POST("/users/registration", userC.CreateUserController)
	e.DELETE("/users", userC.DeleteUserController, echojwt.WithConfig(config))
	e.PUT("/users", userC.UpdateUserController, echojwt.WithConfig(config))
	e.GET("/products/:id", productC.GetProductController, echojwt.WithConfig(config))
	e.GET("/products", productC.GetProductsController, echojwt.WithConfig(config))
	e.GET("/carts/:id", cartC.GetCartController, echojwt.WithConfig(config))
	e.POST("/carts", cartC.CreateCartController, echojwt.WithConfig(config))
	e.DELETE("/carts/:id", cartC.DeleteCartController, echojwt.WithConfig(config))
	e.PUT("/carts/:id", cartC.UpdateCartController, echojwt.WithConfig(config))
	e.GET("/transactions", transactionC.GetTransactionsController, echojwt.WithConfig(config))
	e.GET("/transactions/:id", transactionC.GetTransactionController, echojwt.WithConfig(config))
	e.GET("/transactions/details", transactionDetailC.GetTransactionDetailsController, echojwt.WithConfig(config))
	e.POST("/transactions", transactionC.CreateTransactionController, echojwt.WithConfig(config))

	// development route
	e.PUT("/testransaksi/:id", transactionC.UpdateTransactionController, echojwt.WithConfig(config))

	return e
}
