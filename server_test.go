package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
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
		t.Error("Handler returned wrong status code: \nreceived ",
			status, " \nexpected ",
			http.StatusOK, " \n")
		errors = true
	}

	// Checking for response body difference
	if strings.TrimSpace(rec.Body.String()) != strings.TrimSpace(expectedUser) {
		t.Error("Handler returned unexpected respose body: \nreceived ",
			strings.TrimSpace(rec.Body.String()), " \nexpected ",
			strings.TrimSpace(expectedUser), " \n")
		errors = true
	}

	// If no errors log Success Status
	if !errors {
		log.Println(" GET /users/{id} - getUser PASSED ✅")
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
		t.Error("Handler returned wrong status code: \nreceived ",
			status, " \nexpected ",
			http.StatusOK, " \n")
		errors = true
	}

	// Checking for response body difference
	if strings.TrimSpace(rec.Body.String()) != strings.TrimSpace(expectedBody) {
		t.Error("Handler returned unexpected response body: \nreceived ",
			strings.TrimSpace(rec.Body.String()), "\nexpected ",
			strings.TrimSpace(expectedBody), "{} \n")
		errors = true
	}

	// If no errors log Success Status
	if !errors {
		log.Println(" POST /users/ - createUser PASSED ✅")
	}
}

func TestGetPostHandler(t *testing.T) {
	// Created a new HTTP request to get test post
	r, err := http.NewRequest("GET", "/posts/61682afe8d7c88a454bf269a", nil)
	if err != nil {
		t.Fatal(err)
	}

	// HTTPtest recorder
	rec := httptest.NewRecorder()

	// Retrieving handler to handle posts requests
	handler := http.HandlerFunc(postHandler)

	// String format JSON of expected test post
	expectedPost := `{"_id":"61682afe8d7c88a454bf269a","Caption":"TestPost","ImageURL":"images/test-post.jpg","userid":"61680c9492897f0ebd1fbffa"}`

	// Serving a test HTTP to handle test request
	handler.ServeHTTP(rec, r)

	// Flag variable for errors
	errors := false

	// Checking for status code difference
	if status := rec.Code; status != http.StatusOK {
		t.Error("Handler returned wrong status code: \nreceived ",
			status, " \nexpected ",
			http.StatusOK, " \n")
		errors = true
	}

	// Getting response body
	var resposeBody = strings.TrimSpace(rec.Body.String())

	// Regular Expression to remove timestamp
	reg := regexp.MustCompile(`"timestamp":"([a-zA-Z0-9\- :+=.]+)",`)

	// Removing timestamp from response body
	resposeBody = reg.ReplaceAllString(resposeBody, "")

	// Checking for response body difference
	if strings.TrimSpace(resposeBody) != strings.TrimSpace(expectedPost) {
		t.Error("Handler returned unexpected respose body: \nreceived ",
			strings.TrimSpace(resposeBody), "\nexpected ",
			strings.TrimSpace(expectedPost), "\n")
		errors = true
	}

	// If no errors log Success Status
	if !errors {
		log.Println(" GET /posts/{id} - getPost PASSED ✅")
	}
}

func TestCreatePostHandler(t *testing.T) {
	// String format JSON of expected test post
	json := `{
						"Caption":"testPost1",
						"ImageURL":"images/test-post1.jpg",
						"UserID":"6168231505771cdc4aa206d8"
					}`

	// String converted to bytes array
	jsonBytes := []byte(json)

	// Created a new HTTP request to POST test post
	r, err := http.NewRequest("POST", "/posts/", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Setting headers for HTTP request
	r.Header.Set("Content-Type", "application/json")

	// HTTPtest recorder
	rec := httptest.NewRecorder()

	// Retrieving handler to handle posts requests
	handler := http.HandlerFunc(postHandler)

	// String format JSON of expected body after successful creation of post
	expectedBody := `{"success":"Post Upload successful"}`

	// Serving a test HTTP to handle test request
	handler.ServeHTTP(rec, r)

	// Flag variable for errors
	errors := false

	// Checking for status code difference
	if status := rec.Code; status != http.StatusOK {
		t.Error("Handler returned wrong status code: \nreceived ",
			status, " \nexpected ",
			http.StatusOK, " \n")
		errors = true
	}

	// Checking for response body difference
	if strings.TrimSpace(rec.Body.String()) != strings.TrimSpace(expectedBody) {
		t.Error("Handler returned unexpected response body: \nreceived ",
			strings.TrimSpace(rec.Body.String()), "\nexpected ",
			strings.TrimSpace(expectedBody), "{} \n")
		errors = true
	}

	// If no errors log Success Status
	if !errors {
		log.Println(" POST /posts/ - createPost PASSED ✅")
	}
}

func TestGetUserPostsHandler(t *testing.T) {
	// Created a new HTTP request to get posts of test user
	r, err := http.NewRequest("GET", "/posts/users/61680c9492897f0ebd1fbffa", nil)
	if err != nil {
		t.Fatal(err)
	}

	// HTTPtest recorder
	rec := httptest.NewRecorder()

	// Retrieving handler to handle posts requests
	handler := http.HandlerFunc(userPostHandler)

	// String format JSON of expected test post
	expectedPosts := `[{"Caption":"TestPost","ImageURL":"images/test-post.jpg","_id":"61682afe8d7c88a454bf269a","userid":"61680c9492897f0ebd1fbffa"},{"Caption":"testPost1","ImageURL":"images/test-post1.jpg","_id":"61683f5b9b7e6571a41ae5a5","userid":"61680c9492897f0ebd1fbffa"}]`

	// Serving a test HTTP to handle test request
	handler.ServeHTTP(rec, r)

	// Flag variable for errors
	errors := false

	// Checking for status code difference
	if status := rec.Code; status != http.StatusOK {
		t.Error("Handler returned wrong status code: \nreceived ",
			status, " \nexpected ",
			http.StatusOK, " \n")
		errors = true
	}

	// Getting response body
	var resposeBody = strings.TrimSpace(rec.Body.String())

	// Regular Expression to remove timestamp
	reg := regexp.MustCompile(`"timestamp":"([a-zA-Z0-9\- :+=.]+)",`)

	// Removing timestamp from response body
	resposeBody = reg.ReplaceAllString(resposeBody, "")

	// Checking for response body difference
	if strings.TrimSpace(resposeBody) != strings.TrimSpace(expectedPosts) {
		t.Error("Handler returned unexpected respose body: \nreceived ",
			strings.TrimSpace(resposeBody), "\nexpected ",
			strings.TrimSpace(expectedPosts), "\n")
		errors = true
	}

	// If no errors log Success Status
	if !errors {
		log.Println(" GET /posts/users/{id} - getUserPosts PASSED ✅")
	}
}
