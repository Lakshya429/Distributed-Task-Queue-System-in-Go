package handler

import (
	"log"
	"net/http"

	"github.com/Lakshya429/distributed-task-queue/internal/models"
	"github.com/Lakshya429/distributed-task-queue/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Signup(c *gin.Context) {
	var user SignupRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid Request"})
	}

	log.Println(user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password %v", err)
	}
	user.Password = string(hashedPassword)
	var newUser = models.User{
		ID:       uint(uuid.New().ID()),
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
	 if err := repository.CreateUser(&newUser); err != nil {
	 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	 }
	c.JSON(http.StatusOK, gin.H{"message": "User Created Successfully" , "user" : newUser})
}
