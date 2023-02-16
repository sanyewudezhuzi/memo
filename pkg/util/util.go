package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTsecret = []byte("thisisasecretkey")

type Claims struct {
	ID      uint   `json:"id"`
	Account string `json:"account"`
	jwt.StandardClaims
}

// 生成 token
func GenerateToken(id uint, account string) (string, error) {
	issTime := time.Now()
	expTime := issTime.Add(time.Hour * 24)
	claims := Claims{
		ID:      id,
		Account: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			IssuedAt:  issTime.Unix(),
			Issuer:    "memo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTsecret)
	return tokenString, err
}

// 解析 token
func PatseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
