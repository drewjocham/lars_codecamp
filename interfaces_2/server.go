package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type service struct {
	bService *BookService
}

// example of saving a book with via URL params
//
//localhost:8090/save?name=GoLang&price=12
func (s *service) saveBook(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	name := req.URL.Query().Get("name")
	price := req.URL.Query().Get("price")

	book := &Book{
		Id:    uuid.New(),
		Name:  name,
		Price: price,
	}

	s.bService.SaveBook(ctx, book)
}

// example of getting the request body as an object
/*
curl --location --request POST 'localhost:8090/update' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Id": "c925e9b4-dc7d-4805-8cfe-459bf9f98708",
    "Name": "GoLang",
    "Price": "30"
}'
*/
func (s *service) updateBook(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var b Book
	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.bService.UpdateBook(ctx, &b)
	log.Println("Updated book ", &b)
}

// example gets ID query param in the ?id="1"
// localhost:8090/delete?id=80ec1adf-277b-422c-bb51-de6f39f66166
func (s *service) deleteBook(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	log.Println("Deleting book ", id)
	ctx := context.Background()
	s.bService.DeleteBook(ctx, id)
}

// example of getting {id} in the URL getBook/{id}
// localhost:8090/getBook/"a12a2c02-c226-43da-b500-be014efcb3d3"
func (s *service) getBook(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting book")

	vars := mux.Vars(req)
	id := vars["id"]

	log.Println("id", id)

	book := s.bService.GetBook(id)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}

	return
}

func (s *service) getAllBooks(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	books, err := s.bService.GetAllBooks(ctx)

	if err != nil {
		fmt.Println("[server] error while getting all books")
	}

	fmt.Println(books)
	json.NewEncoder(w).Encode(books)
}

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "demo"
)

func main() {
	log.Println("Starting server")
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)
	postgresDB, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	defer postgresDB.Close()

	repo := NewBookRepository(postgresDB)

	bService := NewBookService(repo)
	s := service{
		bService: bService,
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/save", s.saveBook)
	router.HandleFunc("/update", s.updateBook)
	router.HandleFunc("/delete", s.deleteBook)
	router.HandleFunc("/getBook/{id}", s.getBook)
	router.HandleFunc("/getAllBooks", s.getAllBooks)

	err = http.ListenAndServe(":8090", router)
	if err != nil {
		return
	}

	log.Println("Server started")
}
