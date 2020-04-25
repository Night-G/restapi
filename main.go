package main

import (
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Структура книг (модель)
type Book struct{
	ID		string `json:"id"`
	Isbn	string `json:"isbn"`
	Title	string `json:"title"`
	Author	*Author `json:"Author"`
}
// Структура типа "автор" 
type Author struct{
	FirstName string `json:"firstname"` 
	Lastname  string `json:"lastname"`
}


// получение всех книг 
//Response - ответ / Request - запрос  
func getBooks(w http.ResponseWriter, r *http.Request){

}
// получение одной книги 
func getBook(w http.ResponseWriter, r *http.Request){

}
// создание новой книги 
func createBook(w http.ResponseWriter, r *http.Request){

}
//  обновление книги
func updateBooks(w http.ResponseWriter, r *http.Request){

}
// удаление книги
func deleteBook(w http.ResponseWriter, r *http.Request){

}


func main() {
	router := mux.NewRouter()

	//route handlers /конец
	router.HandleFunc("/api/books", getBooks).Method("GET")
	router.HandleFunc("/api/books/{id}", getBook).Method("GET")
	router.HandleFunc("/api/books", createBook).Method("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Method("PUT")
	router.HandleFunc("/api/books{id}", deleteBook).Method("DELETE")
	log.Fatal(	http.ListenAndServe(":8000",router)	 )
}
