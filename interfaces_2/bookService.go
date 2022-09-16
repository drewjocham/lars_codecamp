package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type BookService struct {
	repo *PostgresRepository
}

type Book struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price string    `json:"price"`
}

func (s *BookService) SaveBook(ctx context.Context, b *Book) {
	fmt.Println("Saving book with id ", b.Id)
	s.repo.save(ctx, b)
}

func (s *BookService) DeleteBook(b *Book) {
	s.repo.delete(b)
}

func (s *BookService) UpdateBook(b *Book) {
	s.repo.update(b)
}

func (s *BookService) GetBook(id string) Book {
	book := s.repo.get(id)
	return book
}

func NewBookService(repo *PostgresRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}
