package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type BookService struct {
	repo *PostgresRepository
}

type Book struct {
	Id    uuid.UUID
	Name  string
	Price string
}

func (s *BookService) SaveBook(ctx context.Context, b *Book) {
	log.Println("[Service] Saving book with id ", b.Id)
	err := s.repo.save(ctx, b)
	if err != nil {
		log.Println("[Service] Error while saving book", err)
	}
}

func (s *BookService) DeleteBook(ctx context.Context, id string) {
	log.Println("[Service] Deleting book with id ", id)
	s.repo.delete(ctx, id)
}

func (s *BookService) UpdateBook(ctx context.Context, b *Book) Book {
	log.Println("[Service] Updating book with id ", b.Id)
	book, err := s.repo.update(ctx, b)
	if err != nil {
		log.Println("[Service] Error while getting book", err)
	}
	return book
}

func (s *BookService) GetBook(id string) Book {
	log.Println("[Service] Getting book with id ", id)
	book, err := s.repo.get(id)
	if err != nil {
		return book
	}
	log.Println("[Service] Error while getting book", err)
	return book
}

func (s *BookService) GetAllBooks() ([]Book, error) {
	log.Println("[Service] Getting books")
	books, err := s.repo.getAllBooks()
	if err != nil {
		fmt.Println("[service] error while getting all books", err)
	}
	return books, nil
}

func NewBookService(repo *PostgresRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}
