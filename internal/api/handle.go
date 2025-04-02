package api

import (
	"github.com/Lakshya429/distributed-task-queue/internal/api/handler"
	"github.com/Lakshya429/distributed-task-queue/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func RoutesHandles(router *gin.Engine) {
	router.POST("/signiup", handler.Signup)
	router.POST("/login", handler.Login)

	authorized := router.Group("/", middleware.AuthMiddleware)
	authorized.POST("/upload", handler.VideoUploadHandler)
	authorized.POST("/serve/:filename", handler.ServeVideo)
}
