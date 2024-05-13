package main

import (
	"log"

	"github.com/BelhajMArouenne1994/GIT_CHAPTER_1/cmd/api_sfmc" // Adjust the import path according to your module's actual path
)

func main() {
	server := api.NewServer()
	err := server.Start(":8080") // Set the port you want to run the server on
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}