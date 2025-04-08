package middleware

import (
    "net/http"
    "strings"

    "backClarineHetic/pkg/jwt" // Notre package d'abstraction
    "github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if strings.HasPrefix(c.Request.URL.Path, "/auth") {
            c.Next()
            return
        }

        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }

        fields := strings.Fields(authHeader)
        if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
            return
        }

        tokenStr := fields[1]

        claims, err := jwt.ValidateToken(tokenStr)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Set("userEmail", claims.Email)
        c.Next()
    }
}
