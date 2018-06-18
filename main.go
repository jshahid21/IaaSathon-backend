package main

import (
	"database/sql"
	"log"
	"os"
)

func main() {
	connectionString := "oracle://"
	connectionString += os.Getenv("ORACLE_USERNAME") + ":"
	connectionString += os.Getenv("ORACLE_PASSWD") + "@"
	connectionString += os.Getenv("ORACLE_SID")
	db, err := sql.Open("goracle", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec("SELECT * FROM *")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(result)
}
