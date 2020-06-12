package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MVC/database"
	"github.com/MVC/database/models"
	"github.com/MVC/database/models/response"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

var db = database.InitDB()

type Exception response.Exception

var JwtVerify = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("akusayangkamu"),
})

func FindUser(username, password string) map[string]interface{} {
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

	errPass := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))

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
