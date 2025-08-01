package auth

import (
	"Accounting/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	TerminalID int64  `json:"terminal_id"`
	ClientID   string `json:"client_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(terminalID int64, clientID string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		TerminalID: terminalID,
		ClientID:   clientID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "accounting",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.JWTSecret))
}
