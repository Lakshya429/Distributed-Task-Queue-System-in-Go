package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"github.com/Lakshya429/distributed-task-queue/internal/kafka"
	"github.com/gin-gonic/gin"
)

const uploadDir = "storage/videos"

func init() {
	// Create the storage directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create storage directory: %v", err)
	}
}

type VideoEvent struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}
var producer *kafka.Producer

func InitProducer(p *kafka.Producer) {
	producer = p
}

func VideoUploadHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatalf("Failed to create storage directory: %v", err)
	}
	defer file.Close()
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)
	filePath := filepath.Join(uploadDir, fileName)

	outFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to saveffds file"})
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file"})
		return
	}
	event := VideoEvent{
		FileName: fileName,
		FilePath: filePath,
	}
	jsonData , err := json.Marshal(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal event"})
		return
	}
	err = producer.Publish("video_events", fileName, string(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File Uploaded Successfully" , "filename" : fileName})
}
