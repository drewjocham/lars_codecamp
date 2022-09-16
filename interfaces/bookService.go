package main

import "github.com/google/uuid"

type BookService struct {
	repo *PostgresRepository
}

type Book struct {
	id    uuid.UUID
	name  string
	price string
}

func (s *BookService) SaveBook(b *Book) {
	s.repo.save(b)
}

func (s *BookService) DeleteBook(b *Book) {
	s.repo.delete(b)
}

func (s *BookService) UpdateBook(b *Book) {
	s.repo.update(b)
}

func NewBookService(repo *PostgresRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}
