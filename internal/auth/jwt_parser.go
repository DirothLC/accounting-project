package auth

import (
	"Accounting/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(getJWTKey())

func getJWTKey() string {
	return config.Cfg.JWTSecret
}

func ParseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
