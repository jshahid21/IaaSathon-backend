package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	_ "gopkg.in/goracle.v2"
)

func main() {

	// Create new router to handle routes
	router := mux.NewRouter()

	router.HandleFunc("/getPolls", GetPolls).Methods("GET")
	router.HandleFunc("/submitPolls", SubmitPoll).Methods("POST")

	// Configure CORS
	c := cors.AllowAll()
	handler := c.Handler(router)

	//serve router
	fmt.Println("Listening on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

// getDBConnection grab environment variables to open
// connection oracle db
func getDBConnection() (db *sql.DB) {
	//required credentials are username, password, and SID
	connectionString := os.Getenv("ORACLE_USERNAME") + "/"
	connectionString += os.Getenv("ORACLE_PASSWORD") + "@"
	connectionString += os.Getenv("ORACLE_SID")

	//open connection
	db, err := sql.Open("goracle", connectionString)
	if err != nil {
		log.Fatal(err)
		return
	}

	//return connection
	return db
}

// GetPolls return poll results for the number of votes for
// either cats or dogs
func GetPolls(w http.ResponseWriter, r *http.Request) {
	var (
		id   int
		name string
	)

	// grab db connection
	db := getDBConnection()
	defer db.Close()

	// query for the sum of cat and dog votes
	rows, err := db.Query("SELECT sum(dog), sum(cat) FROM poll")
	if err != nil {
		log.Fatalf("getPolls error: %s", err)
	}
	defer rows.Close()

	// loop through rows to grab the sum of votes
	// for cat and dog
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	// initialize json with cat and dog votes

	// serve JSON response

}

// SubmitPoll insert into db the vote for
// either cat or dog
func SubmitPoll(w http.ResponseWriter, r *http.Request) {

	//grab db connection to oracle DB
	db := getDBConnection()
	defer db.Close()

	// unmarshal request body to generic interface
	// return map variable
	reqBodyAsJSON := unmarshalJSON(r)

	// prepare insertion statement
	// grab from request body the vote(cat or dog)
	// and create a new row with the vote for cat or dog
	stmt, err := db.Prepare("INSERT INTO poll(?) VALUES(1)")
	if err != nil {
		log.Fatalf("submitPoll error: %s", err)
	}

	// execute statement to DB
	// declare that the vote is for either cat of dog
	_, err = stmt.Exec(reqBodyAsJSON["name"].(string))
	if err != nil {
		log.Fatal(err)
	}

	// return that response is okay

}

// unmarshalJSON from a generic request body
// create a generic map variable by unmarshaling the request
// JSON body
func unmarshalJSON(r *http.Request) (reqBodyAsJSON map[string]interface{}) {
	// declare a generic interface for the JSON bytes to be
	// unmarshaled to
	var genericInterface interface{}

	//read the bytes from the request body
	reqBodyAsBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal the request body bytes to the interface
	err = json.Unmarshal(reqBodyAsBytes, &genericInterface)
	if err != nil {
		log.Fatal(err)
	}

	// create a map variable from the generic interface
	reqBodyAsJSON = genericInterface.(map[string]interface{})
	return
}
