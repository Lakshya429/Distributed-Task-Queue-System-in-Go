package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"github.com/gin-gonic/gin"
)

const uploadDir = "storage/videos"

func init() {
	// Create the storage directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create storage directory: %v", err)
	}
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

	c.JSON(http.StatusOK, gin.H{"message": "File Uploaded Successfully" , "filename" : fileName})
}
