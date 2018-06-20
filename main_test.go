package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
	getPolls(response, request)

	// verify that response code is 200 OK
	checkResponseCode(t, http.StatusOK, response.Code)

	// verify that the the response body is OK
	checkResponseBody(t, response.Body.String(), `{"status": "OK"}`)
}

// TestSubmitPoll test the POST API endpoint
// for submitPolls function
func TestSubmitPoll(t *testing.T) {
	// create test router

	// initialize call method with POST
	// call router with path /submitPoll

	// initialize the header with content-type
	// as application/json

	// call the endpoint

	// verify that response code is 200 OK

	// verify that the content type is application/json
	// in response header

	// verify that response body is a JSON body
	// with "cat": # and "dog": #
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
