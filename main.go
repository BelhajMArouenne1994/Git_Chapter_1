package main

import (
	"log"

	api_soap_to_rest_sfmc "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/cmd/api_soap_to_rest_sfmc" // Adjust the import path according to your module's actual path
)

func main() {
	server := api_soap_to_rest_sfmc.NewServer()
	err := server.Start(":8080") // Set the port you want to run the server on
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
