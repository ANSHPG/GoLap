package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Phase 1 Testing...")
	http.HandleFunc("/upload", fileUploadHandler)

	fmt.Println("Server runing on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fileUploadHandler(response http.ResponseWriter, clientFile *http.Request) {
	clientFile.ParseMultipartForm(20 << 20) // 20mb will be stored

	file, info, err := clientFile.FormFile("myFile")
	if err != nil {
		http.Error(response, "Error rettrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Fprintf(response, "File name: %s", info.Filename)
	fmt.Fprintf(response, "FIle size: %d", info.Size)
	fmt.Fprintf(response, "MIME header: %v", info.Header) //Multipurpose Internet Mail Extension

	location, err := createLocation(info.Filename)
	if err != nil {
		http.Error(response, "Error in saving file", http.StatusInternalServerError)
		return
	}
	defer location.Close()

	_, err = location.ReadFrom(file)
	if err != nil {
		http.Error(response, "Error in saving file", http.StatusInternalServerError)
	}
}

func createLocation(fileName string) (*os.File, error) {
	_, err := os.Stat("uploads")
	if os.IsNotExist(err) {
		os.Mkdir("uploads", 0755) // owner can read/write, others can read
	}

	loaction, err := os.Create(filepath.Join("uploads", fileName))
	if err != nil {
		return nil, err
	}

	return loaction, nil
}
