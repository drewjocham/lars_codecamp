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
	Id    uuid.UUID
	Name  string
	Price string
}

func (s *BookService) SaveBook(ctx context.Context, b *Book) {
	fmt.Println("[Service] Saving book with id ", b.Id)
	s.repo.save(ctx, b)
}

func (s *BookService) DeleteBook(ctx context.Context, id string) {
	fmt.Println("[Service] Deleting book with id ", id)
	s.repo.delete(ctx, id)
}

func (s *BookService) UpdateBook(ctx context.Context, b *Book) {
	fmt.Println("[Service] Updating book with id ", b.Id)
	s.repo.update(ctx, b)
}

func (s *BookService) GetBook(id string) Book {
	fmt.Println("[Service] Getting book with id ", id)
	book := s.repo.get(id)
	return book
}

func NewBookService(repo *PostgresRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}
