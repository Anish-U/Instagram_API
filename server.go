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

	"go.mongodb.org/mongo-driver/bson"
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

type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string             `json:"Caption,omitempty" bson:"Caption,omitempty"`
	ImageURL  string             `json:"ImageURL,omitempty" bson:"ImageURL,omitempty"`
	Timestamp string             `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	UserID    string             `json:"userid,omitempty" bson:"userid,omitempty"`
}

// Handler to handle user related requests
func userHandler(w http.ResponseWriter, r *http.Request) {
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
	userId, _ := primitive.ObjectIDFromHex(strings.TrimPrefix(r.URL.Path, "/users/"))
	
	// A user structure variable
	var user User
	
	// Connecting to Mongo
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	
	// Retrieving Users collection
	collection := client.Database("Instagram").Collection("Users")
	
	// Search Query for Record with ID == userId and if found
	// decode it and store in user structure variable
	err := collection.FindOne(ctx, User{ID: userId}).Decode(&user)

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
func createUser(w http.ResponseWriter, r *http.Request) {
	// Setting the headers to accept JSON
	w.Header().Set("content-type", "application/json")

	// A user structure variable
	var user User

	// Converting the JSON to structure variable 
	_ = json.NewDecoder(r.Body).Decode(&user)
	
	// Initalising hashing
	hash := sha256.New()
	
	// Hash the password using SHA256 and updating int structure variable
	user.Password = (base64.StdEncoding.EncodeToString(hash.Sum([]byte(user.Password))))
	
	// Retrieving Users collection
	collection := client.Database("Instagram").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	
	// Inserting the record into collection
	collection.InsertOne(ctx, user)

	json.NewEncoder(w).Encode(map[string]string{"success": "User Added successful"})
}

// Handler to handle post related requests
func postHandler(w http.ResponseWriter, r *http.Request) {
	// Checking for the method of request
	switch r.Method {
		case http.MethodGet:
			getPost(w, r) 		// directing to getPost handler 
		case http.MethodPost:
			createPost(w, r)	// directing to createPost handler
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	} 
}

// Handler to get specific post by ID
func getPost(w http.ResponseWriter, r *http.Request) {
	// Setting the headers to accept JSON
	w.Header().Set("content-type", "application/json")
	
	// Getting the ID from request URL
	postId, _ := primitive.ObjectIDFromHex(strings.TrimPrefix(r.URL.Path, "/posts/"))
	
	// A post structure variable
	var post Post
	
	// Connecting to Mongo
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	
	// Retrieving Posts collection
	collection := client.Database("Instagram").Collection("Posts")
	
	// Search Query for Record with ID == postId and if found
	// decode it and store in post structure variable
	err := collection.FindOne(ctx, Post{ID: postId}).Decode(&post)

	// If post not found
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// Converting the post structure variable to JSON
	json.NewEncoder(w).Encode(post)
}

// Handler to create a post by JSON body request
func createPost(w http.ResponseWriter, r *http.Request) {
	// Setting the headers to accept JSON
	w.Header().Set("content-type", "application/json")

	// A post structure variable
	var post Post

	// Converting the JSON to structure variable 
	_ = json.NewDecoder(r.Body).Decode(&post)
	
	// Initializing Timestamp to the post
	post.Timestamp = time.Now().String()
	
	// Retrieving Post collection
	collection := client.Database("Instagram").Collection("Posts")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	
	// Inserting the record into collection
	collection.InsertOne(ctx, post)

	json.NewEncoder(w).Encode(map[string]string{"success": "Post Upload successful"})
}


// Handler to handle posts of a specific user requests
func userPostHandler(w http.ResponseWriter, r *http.Request) {
	// Checking for the method of request
	switch r.Method {
		case http.MethodGet:
			getUserPosts(w, r) 		// directing to getUserPosts handler 
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	} 
}

// Handler to get posts specific to a user by ID
func getUserPosts(w http.ResponseWriter, r *http.Request) {
	// Setting the headers to accept JSON
	w.Header().Set("content-type", "application/json")
	
	// Getting the ID from request URL
	userId := strings.TrimPrefix(r.URL.Path, "/posts/users/")
		
	// Connecting to Mongo
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	
	// Retrieving Posts collection
	collection := client.Database("Instagram").Collection("Posts")
	
	// Finding cursor to the all posts where userid == userId
	filterCur, err := collection.Find(ctx, bson.M{"userid": userId})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// Array of post structure
	var posts []bson.M

	// Storing all the post structures into array
	if err = filterCur.All(ctx, &posts); err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return		
	}

	// Converting the array of post structure variable to JSON
	json.NewEncoder(w).Encode(posts)
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
	http.HandleFunc("/users/", userHandler)

	// Handler to handle requests at /users route
	http.HandleFunc("/posts/", postHandler)
	
	// Handler to handle requests at /posts/users route
	http.HandleFunc("/posts/users/", userPostHandler)
	
	// Listening to server at port 8080
	http.ListenAndServe(":8080", nil)
}