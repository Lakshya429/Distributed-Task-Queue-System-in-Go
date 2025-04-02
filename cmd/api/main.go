package main

import (
	"log"
	"github.com/Lakshya429/distributed-task-queue/internal/api"
	"github.com/Lakshya429/distributed-task-queue/pkg/database"
	"github.com/gin-gonic/gin"
	 "github.com/Lakshya429/distributed-task-queue/internal/repository"
	 "github.com/Lakshya429/distributed-task-queue/internal/queues/producer"
)

func main() {
	router := gin.Default()
	database.InitDB()
	repository.Setup()

	_, err :=producer.ConnectionProducer()
	if err != nil {
		log.Fatal("Failed to get Conjumer Channel ", err)
	}

	api.RoutesHandles(router)
	router.Run()
}
