package handler

import (
	"net/http"
	"time"
	"os"
	"github.com/Lakshya429/distributed-task-queue/internal/models"
	"github.com/Lakshya429/distributed-task-queue/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.GetUserbyUserName(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username" : user.Username,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString , err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"error" : err.Error()})
	}
	c.JSON(http.StatusOK , gin.H{"token" : tokenString} )
}
