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
	router.HandleFunc("/books/create", createBook).Methods("POST")        // !N.W!
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")           // !N.W!
	router.HandleFunc("/books/delete/{id}", deleteBook).Methods("DELETE") // !N.W!

	log.Fatal(http.ListenAndServe(":8000", router))
}

//Get All Books
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

//Get Single Book
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
		//books = append(books, book)
	}

	json.NewEncoder(w).Encode(book)
}

//Create a New Book !N.W!
func createBook(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare("INSERT INTO books(id,title) VALUES(@ID,@TITLE)") // incorrect syntax near '?'
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]

	_, err = stmt.Exec(title)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New book was created")
}

//updateBook !N.W!
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := db.Prepare("UPDATE books SET title = @TITLE WHERE id = @ID") // incorrect syntax near '?'
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTitle := keyVal["title"]

	_, err = stmt.Exec(newTitle, params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Book with id = %s was updated", params["id"])
}

//deleteBook   !N.W!
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE from books WHERE id = @ID")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Book with ID %s was deleted", params["id"])
}
