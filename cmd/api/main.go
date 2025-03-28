package main

import (
	"github.com/Lakshya429/distributed-task-queue/internal/api"
	"github.com/Lakshya429/distributed-task-queue/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.InitDB()
	api.RoutesHandles(router)
	router.Run()
}
