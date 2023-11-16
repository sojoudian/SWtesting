package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Student represents the structure for storing student data
type Student struct {
	Name            string `json:"name"`
	ID              string `json:"id"`
	IPAddress       string `json:"ip_address"`
	OperatingSystem string `json:"operating_system"`
	// Add other fields as needed
}

// ConnectMongo creates a connection to the MongoDB
func ConnectMongo() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://0.tcp.ngrok.io:14753")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("StudentDB").Collection("studentRecords")
	return collection
}

// receiveStudentData handles the incoming HTTP requests and stores data in MongoDB
func receiveStudentData(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student.IPAddress = r.RemoteAddr // This gets the IP address
	// student.OperatingSystem = // Extract OS from user-agent or another method

	collection := ConnectMongo()
	_, err = collection.InsertOne(context.TODO(), student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", receiveStudentData).Methods("POST")
	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":9091", router))
}
