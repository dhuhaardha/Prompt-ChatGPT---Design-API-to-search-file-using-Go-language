// // // package main

// // // import (
// // // 	"encoding/json"
// // // 	"flag"
// // // 	"fmt"
// // // 	"net/http"
// // // 	"os"
// // // 	"path/filepath"
// // // )

// // // func searchFilesHandler(w http.ResponseWriter, r *http.Request) {
// // // 	// Parse query parameters
// // // 	query := r.URL.Query()
// // // 	directory := query.Get("directory")
// // // 	fileName := query.Get("filename")

// // // 	if directory == "" || fileName == "" {
// // // 		http.Error(w, "Both 'directory' and 'filename' parameters are required", http.StatusBadRequest)
// // // 		return
// // // 	}

// // // 	// Search for files in the specified directory
// // // 	files, err := searchFiles(directory, fileName)
// // // 	if err != nil {
// // // 		http.Error(w, fmt.Sprintf("Error searching files: %s", err), http.StatusInternalServerError)
// // // 		return
// // // 	}

// // // 	// Return the list of matching files as JSON
// // // 	w.Header().Set("Content-Type", "application/json")
// // // 	json.NewEncoder(w).Encode(files)
// // // }

// // // func searchFiles(directory, fileName string) ([]string, error) {
// // // 	var matchingFiles []string

// // // 	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
// // // 		if err != nil {
// // // 			return err
// // // 		}

// // // 		if !info.IsDir() && info.Name() == fileName {
// // // 			matchingFiles = append(matchingFiles, path)
// // // 		}

// // // 		return nil
// // // 	})

// // // 	return matchingFiles, err
// // // }

// // // func main() {
// // // 	// Parse command line arguments
// // // 	var port int
// // // 	var directory string

// // // 	flag.IntVar(&port, "port", 8080, "Port to run the API server on")
// // // 	flag.StringVar(&directory, "directory", "", "Directory to search for files")
// // // 	flag.Parse()

// // // 	if directory == "" {
// // // 		fmt.Println("Please provide a directory using the -directory flag.")
// // // 		return
// // // 	}

// // // 	// Register the searchFilesHandler function as the handler for the /search endpoint
// // // 	http.HandleFunc("/search", searchFilesHandler)

// // // 	// Start the API server
// // // 	addr := fmt.Sprintf(":%d", port)
// // // 	fmt.Printf("Server is running on http://localhost%s\n", addr)
// // // 	err := http.ListenAndServe(addr, nil)
// // // 	if err != nil {
// // // 		fmt.Printf("Error starting server: %s\n", err)
// // // 	}
// // // }

// // package main

// // import (
// // 	"encoding/json"
// // 	"net/http"
// // 	"os"
// // 	"path/filepath"
// // 	"strings"
// // )

// // type SearchResult struct {
// // 	Path string `json:"path"`
// // }

// // func searchHandler(w http.ResponseWriter, r *http.Request) {
// // 	// Extract the search query from the URL query parameters
// // 	query := r.URL.Query().Get("q")

// // 	// If no search query is provided, return a bad request response
// // 	if query == "" {
// // 		http.Error(w, "Missing search query parameter", http.StatusBadRequest)
// // 		return
// // 	}

// // 	// Specify the root directory to start the search
// // 	rootDir := "." // You can replace this with the desired root directory

// // 	// Perform the search
// // 	results, err := searchFiles(rootDir, query)
// // 	if err != nil {
// // 		http.Error(w, "Error searching files", http.StatusInternalServerError)
// // 		return
// // 	}

// // 	// Convert the search results to JSON and send the response
// // 	w.Header().Set("Content-Type", "application/json")
// // 	json.NewEncoder(w).Encode(results)
// // }

// // func searchFiles(root, query string) ([]SearchResult, error) {
// // 	var results []SearchResult

// // 	// Walk through the directory tree starting from the root
// // 	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// // 		// Check for errors during the walk
// // 		if err != nil {
// // 			return err
// // 		}

// // 		// If the item is a file or directory and contains the search query, add it to the results
// // 		if strings.Contains(info.Name(), query) {
// // 			result := SearchResult{Path: path}
// // 			results = append(results, result)
// // 		}

// // 		return nil
// // 	})

// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	return results, nil
// // }

// // func main() {
// // 	// Define API routes
// // 	http.HandleFunc("/search", searchHandler)

// // 	// Start the server
// // 	port := ":8080"
// // 	http.ListenAndServe(port, nil)
// // }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	directory := "./" // The current directory

// 	files, err := os.Open(directory) //open the directory to read files in the directory
// 	if err != nil {
// 		fmt.Println("error opening directory:", err) //print error if directory is not opened
// 		return
// 	}
// 	defer files.Close() //close the directory opened

// 	fileInfos, err := files.Readdir(-1) //read the files from the directory
// 	if err != nil {
// 		fmt.Println("error reading directory:", err) //if directory is not read properly print error message
// 		return
// 	}
// 	for _, fileInfos := range fileInfos {
// 		fmt.Println(fileInfos.Name()) //print the files from directory
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// type FileResult struct {
// 	Name string `json:"name"`
// 	Size int64  `json:"size"`
// }

// func findFilesByNameHandler(w http.ResponseWriter, r *http.Request) {
// 	// Extract the file name query parameter
// 	fileName := r.URL.Query().Get("name")

// 	// If no file name is provided, return a bad request response
// 	if fileName == "" {
// 		http.Error(w, "Missing file name parameter", http.StatusBadRequest)
// 		return
// 	}

// 	// Specify the root directory to start the search
// 	rootDir := "D:\\"

// 	// Specify the file extension to filter by (e.g., ".zip")
// 	fileExtension := ".zip"

// 	// Perform the search
// 	results, err := findFilesByName(rootDir, fileName, fileExtension)
// 	if err != nil {
// 		http.Error(w, "Error finding files by name", http.StatusInternalServerError)
// 		return
// 	}

// 	// Convert the search results to JSON and send the response
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(results)
// }

// func findFilesByName(root, name, fileExtension string) ([]FileResult, error) {
// 	var results []FileResult

// 	// Walk through the directory tree starting from the root
// 	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// 		// Check for errors during the walk
// 		if err != nil {
// 			return err
// 		}

// 		// If the item is a file, check if it matches the specified name and extension
// 		if !info.IsDir() && strings.HasPrefix(info.Name(), name) && strings.HasSuffix(info.Name(), fileExtension) {
// 			result := FileResult{Name: path, Size: info.Size()}
// 			results = append(results, result)
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return results, nil
// }

// func main() {
// 	// Define API routes
// 	http.HandleFunc("/findByName", findFilesByNameHandler)

// 	// Start the server
// 	port := ":8080"
// 	fmt.Printf("Server is running on port %s...\n", port)
// 	http.ListenAndServe(port, nil)
// }

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
