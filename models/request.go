package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID          uint
	Nickname    string
	AuthorityID uint
	jwt.StandardClaims
}
