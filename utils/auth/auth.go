package auth

import (
	originContext "context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/MVC/database"
	"github.com/MVC/database/models"
	"github.com/MVC/database/models/response"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var db = database.InitDB()

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

func FindUser(username, password string) map[string]interface{} {
	defer db.Close()

	var user models.UsersTable

	err := db.Where("username = ?", username).First(&user).Error

	if err != nil || username != user.Username {
		var res = map[string]interface{}{
			"status":  http.StatusNotFound,
			"message": "Username not found or wrong",
		}

		return res
	}

	expiredAt := time.Now().Add(time.Minute * 100000).Unix()

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if errPass != nil && errPass == bcrypt.ErrMismatchedHashAndPassword {
		var res = map[string]interface{}{
			"status":  http.StatusConflict,
			"message": "Invalid login credentials, Please try again!",
		}

		return res
	}

	tkn := &response.Token{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tkn)

	tokenString, errors := token.SignedString([]byte("akusayangkamu"))

	if errors != nil {
		fmt.Println(errors)
	}

	var res = map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Logged In",
	}
	res["token"] = tokenString
	res["data"] = user

	return res

}
