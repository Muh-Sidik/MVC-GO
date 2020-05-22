package routes

import (
	"net/http"

	"github.com/gocrud/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{"*"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodHead,
			http.MethodOptions,
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
			http.MethodDelete,
		},
	}))

	//route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hai :v")
	})

	e.GET("/user", app.AllUsers)
	e.POST("/user/:name/:email", app.StoreUser)

	//
	return e
}
