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

type polls struct {
	Cat int `json:"cat"`
	Dog int `json:"dog"`
}

type status struct {
	Status string `json:"status"`
}

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
func getDBConnection() (db *sql.DB, err error) {
	//required credentials are username, password, and SID
	connectionString := os.Getenv("ORACLE_USERNAME") + "/"
	connectionString += os.Getenv("ORACLE_PASSWORD") + "@"
	connectionString += os.Getenv("ORACLE_SID")

	//open connection
	db, err = sql.Open("goracle", connectionString)

	//return connection
	return db, err
}

// GetPolls return poll results for the number of votes for
// either cats or dogs
func GetPolls(w http.ResponseWriter, r *http.Request) {

	// grab db connection
	db, err := getDBConnection()
	if err != nil {
		var statusAsJSON status
		statusAsJSON.Status = fmt.Sprintf("Error: %s", err)
		sendBytes(w, statusAsJSON)
		log.Printf("Error with connection to DB: %s\n", err)
		return
	}
	defer db.Close()

	// query for the sum of cat and dog votes
	rows, err := db.Query("SELECT sum(dog), sum(cat) FROM poll")
	if err != nil {
		var statusAsJSON status
		statusAsJSON.Status = fmt.Sprintf("Error: %s", err)
		sendBytes(w, statusAsJSON)
		log.Printf("Error with connection to DB: %s\n", err)
		return
	}
	defer rows.Close()

	// loop through rows to grab the sum of votes
	// for cat and dog
	var pollsAsJSON polls

	for rows.Next() {
		// scan cat and dog results to struct
		err := rows.Scan(&pollsAsJSON.Dog, &pollsAsJSON.Cat)
		if err != nil {
			log.Fatal(err)
		}
	}

	// load struct as bytes to be send
	sendBytes(w, pollsAsJSON)

}

// SubmitPoll insert into db the vote for
// either cat or dog
func SubmitPoll(w http.ResponseWriter, r *http.Request) {

	statusAsJSON := status{
		Status: "ok",
	}

	//grab db connection to oracle DB
	db, err := getDBConnection()
	if err != nil {
		statusAsJSON.Status = fmt.Sprintf("Error: %s", err)
		sendBytes(w, statusAsJSON)
		log.Printf("Error with connection to DB: %s\n", err)
		return
	}
	defer db.Close()

	// unmarshal request body to generic interface
	// return map variable
	reqBodyAsJSON := unmarshalJSON(r)

	// prepare insertion statement
	// grab from request body the vote(cat or dog)
	// and create a new row with the vote for cat or dog
	_, err = db.Exec("INSERT INTO poll(" + reqBodyAsJSON["name"].(string) + ") VALUES(1)")
	if err != nil {
		statusAsJSON.Status = fmt.Sprintf("Error: %s", err)
		sendBytes(w, statusAsJSON)
		log.Printf("submitPoll error: %s\n", err)
		return
	}

	// create struct since status is okay
	sendBytes(w, statusAsJSON)

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

func sendBytes(w http.ResponseWriter, responseAsJSON interface{}) {
	// load struct and convert to bytes to be send
	responseAsBytes, err := json.Marshal(responseAsJSON)
	if err != nil {
		log.Fatal(err)
	}
	// set header and send bytes
	w.Header().Set("Content-type", "application/json")
	w.Write(responseAsBytes)
}
