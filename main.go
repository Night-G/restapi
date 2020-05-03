package main

import (
	"database/sql" //
	//
	//
	//

	//

	_ "github.com/denisenkom/go-mssqldb" //
)

//
//var db *sql.DB
//var err error

func main() {
	//Init Router
	//router := mux.NewRouter()

	//
	db, err := sql.Open("sqlserver", "sqlserver://sa:root@localhost?database=Shop") //
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO Books(BookID, AutorID,title) VALUES (2,3,'t2')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	//Route Handlers/Endpoints
	// router.HandleFunc("/books", getBooks).Methods("GET")
	// router.HandleFunc("/books", createBook).Methods("POST")
	// router.HandleFunc("/books/{id}", getBook).Methods("GET")
	// router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	//log.Fatal(http.ListenAndServe(":8000", router))
}

//Book Struct
// type Book struct {
// 	ID    string `json:"id"`
// 	Title string `json:"title"`
// }

// //books
//var books []Book

//Get All Books
// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	result, err := db.Query("SELECT id, title from books")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer result.Close()

// 	for result.Next() {
// 		var book Book
// 		err := result.Scan(&book.ID, &book.Title)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		books = append(books, book)
// 	}

// 	json.NewEncoder(w).Encode(books)
// }

// //Get Single Book
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) //Get params

// 	result, err := db.Query("SELECT id,title FROM books WHERE id = ?", params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer result.Close()

// 	var book Book

// 	for result.Next() {
// 		var book Book
// 		err := result.Scan(&book.ID, &book.Title)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		books = append(books, book)
// 	}

// 	json.NewEncoder(w).Encode(book)
// }

// //Create a New Book
// func createBook(w http.ResponseWriter, r *http.Request) {
// 	stmt, err := db.Prepare("INSERT INTO books(title) VALUES(?)")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal)
// 	title := keyVal["title"]

// 	_, err = stmt.Exec(title)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Fprintf(w, "New book was created")
// }

// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	stmt, err := db.Prepare("UPDATE books SET title = ? WHERE id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal)
// 	newTitle := keyVal["title"]

// 	_, err = stmt.Exec(newTitle, params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fmt.Fprintf(w, "Book with id = %s was updated", params["id"])
// }

// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	stmt, err := db.Prepare("DELETE from books WHERE id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	_, err = stmt.Exec(params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fmt.Fprintf(w, "Book with ID %s was deleted", params["id"])
// }
