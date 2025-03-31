package main

import (
	"github.com/Lakshya429/distributed-task-queue/internal/api"
	"github.com/gin-gonic/gin"
	
)

func main() {
	router := gin.Default()
	
	api.RoutesHandles(router)
	router.Run()
}
