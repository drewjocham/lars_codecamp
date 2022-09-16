package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type service struct {
	bService *BookService
}

func (s *service) saveBook(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	price := req.URL.Query().Get("price")

	book := &Book{
		id:    uuid.New(),
		name:  name,
		price: price,
	}

	s.bService.SaveBook(book)
}

func (s *service) updateBook(w http.ResponseWriter, req *http.Request) {
	var b Book
	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.bService.UpdateBook(&b)
	fmt.Println("Updated book ", b.name)
}

func (s *service) deleteBook(w http.ResponseWriter, req *http.Request) {
	var b Book
	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.bService.DeleteBook(&b)
}

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "demo"
)

func main() {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)
	postgresDB, err := sql.Open("postgres", connString)

	repo := NewBookRepository(postgresDB)

	bService := NewBookService(repo)
	s := service{
		bService: bService,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/save", s.saveBook)
	mux.HandleFunc("/update/${id}", s.updateBook)
	mux.HandleFunc("/delete/${id}", s.deleteBook)

	err = http.ListenAndServe(":8090", mux)
	if err != nil {
		return
	}
}
