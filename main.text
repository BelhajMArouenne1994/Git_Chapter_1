package main

import (
	"log"

	soap_to_rest_sfmc_api "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/cmd/soap_to_rest_sfmc_api"
)

func main() {
	server := soap_to_rest_sfmc_api.NewServer()
	err := server.Start(":8080") // Set the port you want to run the server on
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
