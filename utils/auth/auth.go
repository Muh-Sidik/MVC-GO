package auth

import (
	originContext "context"
	"net/http"
	"strings"

	"github.com/MVC/database/models/response"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Exception response.Exception

func JwtVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		var header = req.Header.Get("x-access-token")

		header = strings.TrimSpace(header)

		if header == "" {
			//Token missing, 403 unauthorized
			res.WriteHeader(http.StatusForbidden)
			c.JSON(http.StatusForbidden, Exception{Message: "Missing Auth Token"})
		}

		token := &response.Token{}

		_, err := jwt.ParseWithClaims(header, token, func(tk *jwt.Token) (interface{}, error) {
			return []byte("akusayangkamu"), nil
		})

		if err != nil {
			res.WriteHeader(http.StatusForbidden)
			c.JSON(http.StatusForbidden, Exception{Message: err.Error()})
		}

		ctx := originContext.WithValue(req.Context(), "user", token)

		c.Echo().ServeHTTP(res, req.WithContext(ctx))

		return next(c)
	}
}