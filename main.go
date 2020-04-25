package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book struct (модель)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"Author"`
}

//Author struct
type Author struct {
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init Books
var books []Book

// получение всех книг
//Response - ответ / Request - запрос
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// получение одной книги
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) // Get params
	// отбор
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// создание новой книги
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100)) //Mock id - not safe - just for example
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//  обновление книги
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// удаление книги     не работает?
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	//Init Router
	router := mux.NewRouter()

	//Mock data - @todo - implement database
	books = append(books, Book{ID: "1", Isbn: "112233", Title: "FirstBook", Author: &Author{FirstName: "Name0", Lastname: "LastN0"}})
	books = append(books, Book{ID: "2", Isbn: "332211", Title: "SecondBook", Author: &Author{FirstName: "Name1", Lastname: "LastN1"}})
	//books = appen(books, Book{ID: "3", Isbn: "223311", Title: "ThirBook", Author: &Author{FirstName: "Name2", Lastname: "LastN2"}})
	//route handlers /конец
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
