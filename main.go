package main

import (
	"database/sql"
	"log"
	"os"

	_ "gopkg.in/goracle.v2"
)

func main() {
	connectionString := os.Getenv("ORACLE_USERNAME") + "/"
	connectionString += os.Getenv("ORACLE_PASSWD") + "@"
	connectionString += os.Getenv("ORACLE_SID")
	log.Print(connectionString)
	db, err := sql.Open("goracle", connectionString)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %s\n", err)
		return
	} else {
		log.Fatal("Ping Successful")
	}
}

func getPolls(db *sql.DB) {
	result, err := db.Exec("SELECT * FROM poll")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(result)
}

func submitPoll(db *sql.DB) {
	result, err := db.Exec("")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(result)
}
