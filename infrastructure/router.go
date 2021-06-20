package infrastructure

import (
	"net/http"
	"log"
	"app/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/joho/godotenv"
)

func Init() {
	e := echo.New()
	Env_load()

	userController := controllers.NewUserController(NewMySqlDb())
	googleOauth2Controller := controllers.NewGoogleOauth2Controller(NewGetConnect())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/google/oauth2", func(c echo.Context) error { return googleOauth2Controller.Auth(c) })
	e.GET("/google/callback", func(c echo.Context) error { return googleOauth2Controller.Callback(c) })
	e.GET("/users", func(c echo.Context) error { return userController.GetUsers(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.GetUser(c) })
	e.POST("/users", func(c echo.Context) error { return c.String(http.StatusOK, "CreateUser") })
	e.PUT("/users/:id", func(c echo.Context) error { return c.String(http.StatusOK, "UpdateUser") })
	e.DELETE("/users/:id", func(c echo.Context) error { return c.String(http.StatusOK, "DeleteUser") })

	e.Logger.Fatal(e.Start(":8080"))
}

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
