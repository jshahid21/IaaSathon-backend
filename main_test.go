package main

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"
)

// TestMain initialize environment variables from .env file
// in directory
func TestMain(m *testing.M) {

	// open .env file to read and extract
	// environment variables
	// sample env file
	// export ORACLE_USERNAME=<oracle-db-username
	// export ORACLE_PASSWORD=<oracle-db-username-password>
	// export ORACLE_SID=<ip-address>:<db-port>/<unique-db-name>.<service-name>
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create scanner from file to read each line individually
	scanner := bufio.NewScanner(file)

	// create an array where stringArray[0] is the key and
	// stringArray[1] is the value
	var stringArray []string

	// Scan each line in the file
	for scanner.Scan() {
		// Split the line so that stringArray[0] is the keyh
		// and that stringArray[1] is the value
		stringArray = strings.Split(scanner.Text(), "=")

		// set the environment value according to the key in stringArray
		// set pattern for regex to (ORACLE_$$$$) for each string in stringArray[1] to match to
		if re := regexp.MustCompile("(ORACLE_USERNAME)"); re.MatchString(stringArray[0]) {
			os.Setenv("ORACLE_USERNAME", stringArray[1])
		} else if re := regexp.MustCompile("(ORACLE_PASSWORD)"); re.MatchString(stringArray[0]) {
			os.Setenv("ORACLE_PASSWORD", stringArray[1])
		} else if re := regexp.MustCompile("(ORACLE_SID)"); re.MatchString(stringArray[0]) {
			os.Setenv("ORACLE_SID", stringArray[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// debug to verify that environment values were initialized properly
	log.Println("ORACLE_USERNAME: " + os.Getenv("ORACLE_USERNAME"))
	log.Println("ORACLE_PASSWORD: " + os.Getenv("ORACLE_PASSWORD"))
	log.Println("ORACLE_SID: " + os.Getenv("ORACLE_SID"))
}

// TestGetPolls test the GET REST API endpoint
// for getPolls function
func TestGetPolls(t *testing.T) {
	// create test router
	log.Print("testing getting polls")

	request, err := http.NewRequest("GET", "/getPoll", nil)
	if err != nil {
		log.Fatal(err)
	}
	response := httptest.NewRecorder()
	// call router with path /getPolls and GET method

	// verify that content-type is application/json
	request.Header.Add("Content-type", "application/json")

	// call the endpoint
	GetPolls(response, request)

	// verify that response code is 200 OK
	checkResponseCode(t, http.StatusOK, response.Code)

	// verify that the the response body is OK

	// TODO
	// modify actual string for regex pattern
	checkResponseBody(t, response.Body.String(), `{"status": "OK"}`)
}

// TestSubmitPoll test the POST API endpoint
// for submitPolls function
func TestSubmitPoll(t *testing.T) {
	// create test router
	log.Print("testing submitting poll")

	message := []byte(`{"name": "cat"}`)
	// initialize call method with POST
	// call router with path /submitPoll
	request, err := http.NewRequest("POST", "/submitPoll", bytes.NewBuffer(message))
	if err != nil {
		log.Fatal(err)
	}
	response := httptest.NewRecorder()
	// initialize the header with content-type
	// as application/json
	request.Header.Add("Content-type", "application/json")
	// call the endpoint
	SubmitPoll(response, request)

	// verify that response code is 200 OK
	checkResponseCode(t, http.StatusOK, response.Code)

	// verify that response body is a JSON body
	// with "cat": # and "dog": #

	// TODO
	// modify actual string to regex pattern
	checkResponseBody(t, response.Body.String(), `{"cat": [0-9*], "dog": [0-9*]}`)
}

// checkResponseCode verify the response code
func checkResponseCode(t *testing.T, expected, actual int) {
	// if response status is equal to tested status then ok
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// checkResponseBody verify the response body
func checkResponseBody(t *testing.T, actualBody, expectedBody string) {
	// if response body is equal to tested body then ok
	if actualBody != expectedBody {
		t.Errorf("Expected %s. Got %s\n", expectedBody, actualBody)
	}
}
