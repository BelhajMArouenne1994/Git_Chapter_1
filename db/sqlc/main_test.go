package db

import (
	"log"
	"testing"
    "database/sql"
    "os"
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


var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){
	var err error
	testDB, err = sql.Open(dbDriver, "postgresql://" +  user + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable")
	if err!= nil {
        log.Fatal(err)
    }

	testQueries = New(testDB)
	os.Exit(m.Run())
}