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
	s.repo.save(ctx, b)
}

func (s *BookService) DeleteBook(ctx context.Context, id string) {
	log.Println("[Service] Deleting book with id ", id)
	s.repo.delete(ctx, id)
}

func (s *BookService) UpdateBook(ctx context.Context, b *Book) {
	log.Println("[Service] Updating book with id ", b.Id)
	s.repo.update(ctx, b)
}

func (s *BookService) GetBook(id string) Book {
	log.Println("[Service] Getting book with id ", id)
	book := s.repo.get(id)
	return book
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]Book, error) {
	log.Println("[Service] Getting books")
	books, err := s.repo.getAllBooks(ctx)
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
