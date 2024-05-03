package main

import (
	"database/sql"
	"log"

	api "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/api"
	db "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	host     = "localhost"
	port     = "5432"
	user     = "root"
	password = "secret"
	database = "nps-postgres"
)

var testQueries *db.Queries
var databaseCurrent *sql.DB

func main() {

	var err error
	databaseCurrent, err = sql.Open(dbDriver, "postgresql://"+user+":"+password+"@"+host+":"+port+"/"+database+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	nps := db.NewNps(databaseCurrent)

	server := api.NewServer(nps)
	server.Start("0.0.0.0:8080")
}
