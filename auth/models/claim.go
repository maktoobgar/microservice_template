package models

import "github.com/dgrijalva/jwt-go"

var RefreshTokenType = "1"
var AccessTokenType = "2"

type Claims struct {
	ID   int64
	Type string
	jwt.StandardClaims
}
