package main

import (
	"github.com/Lakshya429/distributed-task-queue/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/upload", handler.VideoUploadHandler)
	router.GET("/get/:filename" , handler.ServeVideo)
	router.Run()
}
