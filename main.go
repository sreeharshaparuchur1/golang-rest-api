package main

import (
  //"fmt"
  //"encoding/json"
  "log"
  "net/http"
  //"math/rand"
  //"strconv"
  "github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct {
  ID  string "json:id"
  Isbn  string "json:isbn"
  Title  string "json:title"
  Author  *Author "json:author"
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//slice - variable length array
var books []Book



//Get all books
// Every route handler has to take in a response and a request
func getBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}



func main() {
  // Init Router
  r := mux.NewRouter()

  // Mock Data - can read from a database
  books = append(books, Book{ID: "1", Isbn: "42069". Title: "The Subtle art", Author: &Author{Firstname: "Abel", Lastname: "Tesfaye"}})
  books = append(books, Book{ID: "2", Isbn: "23112". Title: "Not a Penny More, Not a Penny Less", Author: &Author{Firstname: "Jeffery", Lastname: "Archer"}})
  books = append(books, Book{ID: "3", Isbn: "27271". Title: "Digital Fortress", Author: &Author{Firstname: "Dan", Lastname: "Brown"}})
  // Route Handlers / Endpoints
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/books", createBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", r))

}
