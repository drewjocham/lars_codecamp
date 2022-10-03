package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

type service struct {
	bService *BookService
	cache    *BookCache
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
	json.NewEncoder(w).Encode(book)
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
	json.NewEncoder(w).Encode(b)
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

// localhost:8090/getBook?id=80ec1adf-277b-422c-bb51-de6f39f66166
func (s *service) getBook(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting book")

	id := req.URL.Query().Get("id")

	log.Println("id", id)

	//book := s.bService.GetBook(id)
	// Now lets use the cache
	book, ok := s.cache.GetBookById(id)
	if !ok {
		log.Println("Was not able to get book from cache")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (s *service) getAllBooks(w http.ResponseWriter, req *http.Request) {

	//books, err := s.bService.GetAllBooks()
	books := s.cache.GetAllBooks()

	//if err != nil {
	//	fmt.Println("[server] error while getting all books")
	//}

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
	bookCache := NewBookCache(repo)

	bService := NewBookService(repo)
	s := service{
		bService: bService,
		cache:    bookCache,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, egCtx := errgroup.WithContext(ctx)

	t := time.Now()
	err = s.cache.Init(5, time.Duration(t.Second()+15))
	if err != nil {
		log.Println("error while loading cache", err)
	}

	g.Go(func() error {
		log.Println("starting campaign cache refresh")

		return s.cache.StartRefresh(egCtx)
	})

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/save", s.saveBook)
	router.HandleFunc("/update", s.updateBook)
	router.HandleFunc("/delete", s.deleteBook)
	router.HandleFunc("/getBook", s.getBook)
	router.HandleFunc("/getAllBooks", s.getAllBooks)

	err = http.ListenAndServe(":8090", router)
	if err != nil {
		return
	}

	log.Println("Server started")
}
