package mongo_db_api

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    mongo_db "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db"

	
    //handlers "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db_handlers"
)

func main() {
    // Connexion à la base de données
    client, ctx, cancel := mongo_db.ConnectDB()
    defer cancel()
    defer client.Disconnect(ctx)

    // Initialisation du routeur
    router := mux.NewRouter()

    // Routes
    //router.HandleFunc("/dataextensions", handlers.GetDataExtensions).Methods("GET")
    //router.HandleFunc("/dataextensions/{id}", handlers.GetDataExtension).Methods("GET")
    //router.HandleFunc("/dataextensions", handlers.CreateDataExtension).Methods("POST")
    //router.HandleFunc("/dataextensions/{id}", handlers.UpdateDataExtension).Methods("PUT")
    //router.HandleFunc("/dataextensions/{id}", handlers.DeleteDataExtension).Methods("DELETE")

    // Démarrage du serveur
    log.Fatal(http.ListenAndServe(":8080", router))
}
