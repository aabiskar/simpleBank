package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/aabiskar/simplebank/util"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	log.Printf("Loaded config for test: %+v", config)

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	//
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	if err := testDB.Ping(); err != nil {
		log.Fatal("cannot ping db:", err)
	}

	log.Println("connected to db")

	testQueries = New(testDB)

	os.Exit(m.Run())
}
