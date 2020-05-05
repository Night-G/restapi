package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

//Book Struct
type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

//books
var books []Book

func main() {
	//init router
	router := mux.NewRouter()

	db, err = sql.Open("sqlserver", "sqlserver://sa:root@localhost?database=Shop")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//route endpoints
	router.HandleFunc("/books", getBooks).Methods("GET")                  // works
	router.HandleFunc("/books/{id}", getBook).Methods("GET")              // works
	router.HandleFunc("/books/create", createBook).Methods("POST")        // works
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")           // works
	router.HandleFunc("/books/delete/{id}", deleteBook).Methods("DELETE") // works

	log.Fatal(http.ListenAndServe(":8000", router))
}

//Get All Books ../books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book

	result, err := db.Query("SELECT id, title from books")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var book Book
		err := result.Scan(&book.ID, &book.Title)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

//Get Single Book ../books/{id}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	result, err := db.Query("SELECT * FROM books WHERE id = @ID ", sql.Named("id", params["id"])) // incorrect syntax near '?' !исправлено!
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var book Book

	for result.Next() {
		err := result.Scan(&book.ID, &book.Title)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(book)
}

//Create a New Book ../books/create/{id}
func createBook(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]
	id := keyVal["id"]

	a, err := db.Query("INSERT INTO books VALUES(@id,@title)", sql.Named("id", id), sql.Named("title", title))
	if err != nil {
		panic(err.Error())
	}

	defer a.Close()

	fmt.Fprintf(w, "New book was created")
}

//updateBook ../books/{id}
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTitle := keyVal["title"]

	ad, err := db.Query("UPDATE Books SET Title = @newTitle WHERE id = @id",
		sql.Named("newTitle", newTitle), sql.Named("id", params["id"]))
	if err != nil {
		panic(err.Error())
	}
	defer ad.Close()
	fmt.Fprintf(w, "Book with ID = %s was updated", params["ID"])
}

//deleteBook ../books/delete/{id}
func deleteBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) //Get params

	result, err := db.Query("DELETE FROM books WHERE id = @ID ", sql.Named("id", params["id"]))
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	fmt.Fprintf(w, "Book with ID %s was deleted", params["id"])
}
