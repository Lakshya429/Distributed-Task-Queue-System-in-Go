package main

import (
	"github.com/Lakshya429/distributed-task-queue/internal/api"
	"github.com/Lakshya429/distributed-task-queue/pkg/database"
	"github.com/gin-gonic/gin"
	 "github.com/Lakshya429/distributed-task-queue/internal/repository"
)

func main() {
	router := gin.Default()
	database.InitDB()
	repository.Setup()
	api.RoutesHandles(router)
	router.Run()
}
