package infrastructure

import (
	"app/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func Init() {
	e := echo.New()

	userController := controllers.NewUserController(NewMySqlDb())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(c echo.Context) error { return userController.GetUsers(c) })
	e.GET("/users/:id", func(c echo.Context) error { return c.String(http.StatusOK, "GetUser") })
	e.POST("/users", func(c echo.Context) error { return c.String(http.StatusOK, "CreateUser") })
	e.PUT("/users/:id", func(c echo.Context) error { return c.String(http.StatusOK, "UpdateUser") })
	e.DELETE("/users/:id", func(c echo.Context) error { return c.String(http.StatusOK, "DeleteUser") })

	e.Logger.Fatal(e.Start(":8080"))
}
