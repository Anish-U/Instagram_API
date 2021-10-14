package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetUserHandler(t *testing.T) {

	// Created a new HTTP request to get test user
	r, err := http.NewRequest("GET", "/users/61680c9492897f0ebd1fbffa", nil)
	if err != nil {
		t.Fatal(err)
	}

	// HTTPtest recorder
	rec := httptest.NewRecorder()

	// Retrieving handler to handle users requests
	handler := http.HandlerFunc(userHandler)

	// String format JSON of expected test user
	expectedUser := `{"_id":"61680c9492897f0ebd1fbffa","Name":"testUser","email":"testEmail@mail.com","password":"dGVzdFBhc3N3b3Jk47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU="}`

	// Serving a test HTTP to handle test request
	handler.ServeHTTP(rec, r)

	// Flag variable for errors
	errors := false

	// Checking for status code difference
	if status := rec.Code; status != http.StatusOK {
		t.Error("Handler returned wrong status code: \nreceived {} \nexpected {} \n",
			status, http.StatusOK)
		errors = true
	}

	// Checking for response body difference
	if strings.TrimSpace(rec.Body.String()) != strings.TrimSpace(expectedUser) {
		t.Error("Handler returned unexpected respose body: \nreceived {} \nexpected {} \n",
			strings.TrimSpace(rec.Body.String()),
			strings.TrimSpace(expectedUser))
		errors = true
	}

	// If no errors log Success Status
	if !errors {
		log.Println(" GET /users/{id} - getUsers PASSED ✅")
	}
}

func TestCreateUserHandler(t *testing.T) {

	// String format JSON of expected test user
	json := `{
						"Name":"testUser1",
						"Email":"testUser1@mail.com",
						"Password":"testPassword1"
					}`

	// String converted to bytes array
	jsonBytes := []byte(json)

	// Created a new HTTP request to post test user
	r, err := http.NewRequest("POST", "/users/", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Setting headers for HTTP request
	r.Header.Set("Content-Type", "application/json")

	// HTTPtest recorder
	rec := httptest.NewRecorder()

	// Retrieving handler to handle users requests
	handler := http.HandlerFunc(userHandler)

	// String format JSON of expected body after successful creation of user
	expectedBody := `{"success":"User Added successful"}`

	// Serving a test HTTP to handle test request
	handler.ServeHTTP(rec, r)

	// Flag variable for errors
	errors := false

	// Checking for status code difference
	if status := rec.Code; status != http.StatusOK {
		t.Error("Handler returned wrong status code: \nreceived {} \nexpected {} \n",
			status, http.StatusOK)
		errors = true
	}

	// Checking for response body difference
	if strings.TrimSpace(rec.Body.String()) != strings.TrimSpace(expectedBody) {
		t.Error("Handler returned unexpected response body: \nreceived {} \nexpected {} \n",
			strings.TrimSpace(rec.Body.String()),
			strings.TrimSpace(expectedBody))
		errors = true
	}

	// If no errors log Success Status
	if !errors {
		log.Println(" POST /users/ - createUser PASSED ✅")
	}
}
