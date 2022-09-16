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
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			// log some error
		}
	}(tx)

	_, err = tx.Exec(`INSERT INTO Book(name, price) VALUES ($1, $2)`, s.name, s.price)
	if err != nil {
		return
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			// log some error
		}
	}(r.DB)
}

func (r *PostgresRepository) delete(s *Book) {

	tx, err := r.DB.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			// log some error
		}
	}(tx)

	_, err = tx.Exec(`DELETE FROM Book WHERE id = ($1)`, s.id)
	if err != nil {
		return
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			// log some error
		}
	}(r.DB)

	log.Println("DELETE")
}

func (r *PostgresRepository) update(s *Book) {

	tx, err := r.DB.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			// log some error
		}
	}(tx)

	_, err = tx.Exec(`UPDATE Book SET name=$1, price=$2 WHERE id = $3`, s.name, s.price, s.id)
	if err != nil {
		return
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			// log some error
		}
	}(r.DB)

	log.Println("UPDATE: Name: " + s.name + " | Price: " + s.price)
}

func NewBookRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		DB: db,
	}
}
