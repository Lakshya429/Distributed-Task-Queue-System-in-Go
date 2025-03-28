package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)
var secretKey = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized , gin.H{"error" : "request does not contain an access token"})
		c.Abort()
		return
	}

	parts := strings.Split(tokenString , " ")

	if len(parts) < 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized , gin.H{"error" : "wrong header format"})
		c.Abort()
		return
	}

	token := parts[1]

	validate , err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	 })

	 if err != nil || !validate.Valid {
		c.JSON(http.StatusUnauthorized , gin.H{"error" : "invalid token"})
		c.Abort()
		return
	 }
	 c.Next()
}