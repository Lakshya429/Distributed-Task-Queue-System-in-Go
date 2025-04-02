package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"fmt"
	"github.com/gin-gonic/gin"
)

var filePath = "C:/Users/cheli/Distributed-Task-Queue-System/cmd/api/storage/videos"

func ServeVideo(c *gin.Context) {
	fileName := c.Param("filename") // Get filename from URL
	filePath := filepath.Join(filePath, fileName)
	fmt.Println(filePath , fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	// Serve the file
	c.File(filePath)
}
