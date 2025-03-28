package handler

import (
	"log"
	"net/http"

	"github.com/Lakshya429/distributed-task-queue/internal/models"
	"github.com/Lakshya429/distributed-task-queue/internal/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c * gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
	}

	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password %v" , err)
	}
	user.Password = string(hashedPassword)
	if err := repository.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"error" : err.Error()})
	}
	c.JSON(http.StatusOK , gin.H{"message" : "User Created Successfully"})
}