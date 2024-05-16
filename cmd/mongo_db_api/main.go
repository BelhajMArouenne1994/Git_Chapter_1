package mongo_db_api

import (
    //"log"
    //"net/http"
    //"github.com/gorilla/mux"
    "github.com/gin-gonic/gin"


    mongo_db "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db"

	
    handlers "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db_handlers"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
    // Connexion à la base de données
    client, ctx, cancel := mongo_db.ConnectDB()
    defer cancel()
    defer client.Disconnect(ctx)

    // Initialisation du routeur
 	server := Server{}

	router := gin.New()

    // Routes
    //router.HandleFunc("/dataextensions", handlers.GetDataExtensions).Methods("GET")
    //router.HandleFunc("/dataextensions/{id}", handlers.GetDataExtension).Methods("GET")
    //router.HandleFunc("/dataextensions/{id}", handlers.UpdateDataExtension).Methods("PUT")
    //router.HandleFunc("/dataextensions/{id}", handlers.DeleteDataExtension).Methods("DELETE")




    router.POST("/dataextensions", handlers.Handler(handlers.CreateDataExtension()))



	server.router = router
	return &server
}
