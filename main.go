package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func main(){

	r:= mux.NewRouter()

	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("api/users", createUser).Methods("Post")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getUsers(w http.ResponseWriter, r *http.Request){
	users := []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}

func getUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userID := vars["id"]
	w.Write([]byte("Getting user: " + userID))
}

func updateUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("User updated"))
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("User deleted"))
}