package main

import "testing"

// TestGetPolls test the GET REST API endpoint
// for getPolls function
func TestGetPolls(t *testing.T) {
	// create test router

	// call router with path /getPolls and GET method

	// verify that content-type is application/json

	// call the endpoint

	// verify that response code is 200 OK

	// verify that the response header
	// content-type is application/json

	// verify that the the response body is OK
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
