package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Lakshya429/distributed-task-queue/internal/models"
)


func Login(c * gin.Context) {
	 var loginRequest models.LoginRequest 

	 if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
		return
	 }

	 


}