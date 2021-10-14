package main

import (
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

	errors := false
	// Checking for status code difference
	if status := rec.Code; status != http.StatusOK {
		t.Error("Handler returned wrong status code: \nreceived {} \nexpected {} \n",
			status, http.StatusOK)
		errors = true
	}

	// Checking for body content difference
	if strings.TrimSpace(rec.Body.String()) != strings.TrimSpace(expectedUser) {
		t.Error("Handler returned unexpected body content: \nreceived {} \nexpected {} \n",
			strings.TrimSpace(rec.Body.String()),
			strings.Trim(expectedUser, "\\s"))
		errors = true
	}

	// If no errors log Success Status
	if !errors {
		log.Println(" GET /users/ - getUsers PASSED ✅")
	}
}
