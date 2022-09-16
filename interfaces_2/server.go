package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

type service struct {
	bService *BookService
}

func (s *service) saveBook(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	name := req.URL.Query().Get("name")
	price := req.URL.Query().Get("price")

	fmt.Println("name", name)
	fmt.Println("price", price)

	book := &Book{
		Id:    uuid.New(),
		Name:  name,
		Price: price,
	}

	s.bService.SaveBook(ctx, book)
}

func (s *service) updateBook(w http.ResponseWriter, req *http.Request) {
	var b Book
	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.bService.UpdateBook(&b)
	fmt.Println("Updated book ", b.Name)
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

func (s *service) getBook(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting book")

	//ctx := context.Background()
	vars := mux.Vars(req)
	id := vars["id"]

	fmt.Println("id", id)

	book := s.bService.GetBook(id)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}

	return
}

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "demo"
)

func main() {
	fmt.Println("Starting server")
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
	router.HandleFunc("/update/{id}", s.updateBook)
	router.HandleFunc("/delete/{id}", s.deleteBook)
	router.HandleFunc("/getBook/{id}", s.getBook)

	err = http.ListenAndServe(":8090", router)
	if err != nil {
		return
	}

	fmt.Println("Server started")
}
