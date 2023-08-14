package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	ID     int     `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

type Author struct {
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
}

var books []Book

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func main2() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", getBooks).Method("GET")
	http.Handle("/", r)

}
