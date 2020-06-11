package routes

import (
	"net/http"

	"github.com/MVC/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "golang.org/x/crypto/acme/autocert"
)

func Init() *echo.Echo {
	e := echo.New()

	//host list
	// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<DOMAIN>")
	// Cache certificates
	// e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

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
	e.POST("/login", app.Login)
	e.POST("/user/:name/:email", app.Register)

	// user authorized
	// e.Use(auth.JwtVerify)
	e.GET("/user", app.AllUsers)
	//
	return e
}
