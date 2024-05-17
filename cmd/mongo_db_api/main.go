package main

import (
	//"context"
	"log"

	"github.com/gin-gonic/gin"
	client "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db/client"
	handlers "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db/handlers"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	// Initialisation du routeur
	server := Server{}
	router := gin.New()

	// Connexion à la base de données
	mongoClient, ctx, cancel := client.ConnectDB()
	defer func() {
		cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("Error disconnecting from database: %v", err)
		}
	}()

	// Passer la connexion de la base de données au gestionnaire
	// Les gestionnaires peuvent utiliser cette connexion via le contexte
	router.Use(func(c *gin.Context) {
		c.Set("dbClient", mongoClient)
		c.Set("dbContext", ctx)
		c.Next()
	})

	// Routes
	router.POST("/dataextensions", handlers.CreateDataExtension())

	server.router = router
	return &server
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func main() {
	server := NewServer()
	if err := server.Start(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
