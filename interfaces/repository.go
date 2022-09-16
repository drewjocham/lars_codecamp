package main

import (
	"database/sql"
	"log"
)

type PostgresRepository struct {
	DB *sql.DB
}

func (r *PostgresRepository) save(s *Book) {
	tx, err := r.DB.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	tx.Exec(`INSERT INTO Book(name, price) VALUES ($1, $2)`, s.name, s.price)

	defer r.DB.Close()
}

func (r *PostgresRepository) delete(s *Book) {

	tx, err := r.DB.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	tx.Exec(`DELETE FROM Book WHERE id = ($1)`, s.id)

	log.Println("DELETE")
}

func (r *PostgresRepository) update(s *Book) {

	tx, err := r.DB.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	tx.Exec(`UPDATE Book SET name=$1, price=$2 WHERE id = $3`, s.name, s.price, s.id)

	log.Println("UPDATE: Name: " + s.name + " | Price: " + s.price)
}

func NewBookRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		DB: db,
	}
}
