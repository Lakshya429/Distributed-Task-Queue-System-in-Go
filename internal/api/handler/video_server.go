package handler 


import (
	"net/http"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
)


func ServeVideo(c *gin.Context) {
	fileName := c.Param("filename") // Get filename from URL
	filePath := filepath.Join(uploadDir, fileName)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Serve the file
	c.File(filePath)
}