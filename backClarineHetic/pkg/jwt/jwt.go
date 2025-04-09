package jwt

import (
    jwtLib "github.com/golang-jwt/jwt/v4"
    "time"
)

type CustomClaims struct {
    Email    string `json:"email"`
    UserUUID string `json:"user_uuid"`
    jwtLib.RegisteredClaims
}

var jwtKey = []byte("secret_key")

func GenerateToken(email string, userUUID string) (string, error) {
    claims := &CustomClaims{
        Email:    email,
        UserUUID: userUUID,
        RegisteredClaims: jwtLib.RegisteredClaims{
            Subject:   email,
            ExpiresAt: jwtLib.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwtLib.NewNumericDate(time.Now()),
        },
    }
    token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*CustomClaims, error) {
    token, err := jwtLib.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwtLib.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, err
}
