package main

// Importing dependencies
import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo Client
var client *mongo.Client

// Model for User Collection
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

// Handler to handle user related requests
func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Checking for the method of request
	switch r.Method {
		case http.MethodGet:
			getUser(w, r) 		// directing to getUser handler 
		case http.MethodPost:
			createUser(w, r)	// directing to createUser handler
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	} 
}

// Handler to get specific user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	// Setting the headers to accept JSON
	w.Header().Set("content-type", "application/json")
	
	// Getting the ID from request URL
	id, _ := primitive.ObjectIDFromHex(strings.TrimPrefix(r.URL.Path, "/users/"))
	
	// A user structure variable
	var user User
	
	// Connecting to Mongo
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	
	// Retrieving Users collection
	collection := client.Database("Instagram").Collection("Users")
	
	// Search Query for Record with ID == id and if found
	// decode it and store in user structure variable
	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)

	// If user not found
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// Converting the user structure variable to JSON
	json.NewEncoder(w).Encode(user)
}

// Handler to create a user by JSON body request
func createUser(res http.ResponseWriter, req *http.Request) {
	// Setting the headers to accept JSON
	res.Header().Set("content-type", "application/json")

	// A user structure variable
	var user User

	// Converting the JSON to structure variable 
	_ = json.NewDecoder(req.Body).Decode(&user)
	
	// Initalising hashing
	hash := sha256.New()
	
	// Hash the password using SHA256 and updating int structure variable
	user.Password = (base64.StdEncoding.EncodeToString(hash.Sum([]byte(user.Password))))
	
	// Retrieving Users collection
	collection := client.Database("Instagram").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	
	// Inserting the record into collection
	collection.InsertOne(ctx, user)

	json.NewEncoder(res).Encode(map[string]string{"success": "Upload successful"})
}

// Main Method
func main() {
	
	fmt.Println("Server Running at http://localhost:8080")

	// Connecting to Mongo
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	// Handler to handle requests at /users route
	http.HandleFunc("/users/", UserHandler)
	
	// Listening to server at port 8080
	http.ListenAndServe(":8080", nil)
}