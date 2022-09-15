package main

import "github.com/google/uuid"

type repository interface {
	save(s Book)
	delete(s Book)
	update(s Book)
}

type BookService struct {
	r *repository
}

type Book struct {
	id    uuid.UUID
	name  string
	price float64
}

func (s *BookService) saveBook(b *Book) {
	s.saveBook(b)
}

func (s *BookService) deleteBook(b *Book) {
	s.deleteBook(b)
}

func (s *BookService) updateBook(b *Book) {
	s.updateBook(b)
}
