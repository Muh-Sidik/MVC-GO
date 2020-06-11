package response

import jwt "github.com/dgrijalva/jwt-go"

type Token struct {
	ID       uint
	Username string
	Password string

	*jwt.StandardClaims
}
