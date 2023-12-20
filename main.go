package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/files", findFile)

	r.Run(":8080")
}

func findFile(c *gin.Context) {
	// Get the filename from the query parameter
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing filename parameter"})
		return
	}

	// Specify the directory (in this case, "D:")
	directory := "D:\\"

	// Search for the file in the specified directory
	result, err := findFileInDirectory(directory, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "File not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func findFileInDirectory(directory, filename string) (string, error) {
	var result string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current file matches the desired filename
		if !info.IsDir() && info.Name() == filename {
			result = path
			return fmt.Errorf("file found") // Stop walking the directory once the file is found
		}

		return nil
	})

	// If the file is not found, err will be set to "file not found"
	if err != nil && err.Error() != "file found" {
		return "", err
	}

	return result, nil
}
