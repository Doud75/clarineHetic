package jwt

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("secret_key")

func GenerateToken(email string) (string, error) {
    claims := &jwt.RegisteredClaims{
        Subject:   email,
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}
