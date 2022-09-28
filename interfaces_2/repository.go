package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

var (
	ErrBookNotFound = errors.New("book does not exist")
)

type PostgresRepository struct {
	DB *sql.DB
}

func (r *PostgresRepository) save(ctx context.Context, s *Book) error {
	var tx *sql.Tx

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = r.commitOrRollback(err, tx)
	}()

	_, err = tx.Exec(`INSERT INTO demo.book(id, name, price) VALUES ($1, $2, $3)`, s.Id, s.Name, s.Price)
	if err != nil {
		log.Println("Error when trying to save ", err)
		return ErrBookNotFound
	}

	return nil
}

func (r *PostgresRepository) delete(ctx context.Context, id string) {

	var tx *sql.Tx

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error while deleting book ", id)
	}

	defer func() {
		err = r.commitOrRollback(err, tx)
	}()

	var dbRes sql.Result

	dbRes, err = tx.Exec(`DELETE FROM demo.book WHERE id = ($1)`, id)
	if err != nil {
		return
	}

	var rows int64

	rows, err = dbRes.RowsAffected()
	if err != nil {
		log.Println("Error while deleting book row ")
	}

	if rows != 1 {
		log.Println("Book was not found, unable to delete.")
	}

	log.Println("Deleted book")
}

func (r *PostgresRepository) update(ctx context.Context, s *Book) (Book, error) {
	log.Println("[Repo] Updating book with id ", s.Id)
	tx, err := r.DB.BeginTx(ctx, nil)

	if err != nil {
		log.Println("Error in update translation ", err)
	}

	defer func() {
		err = r.commitOrRollback(err, tx)
	}()

	book := Book{}

	_, err = tx.Exec(`UPDATE demo.book SET name=$1, price=$2 WHERE id = $3`, s.Name, s.Price, s.Id)
	if err != nil {
		return book, err
	}

	row := r.DB.QueryRow(`SELECT * FROM demo.book WHERE id = $1`, s.Id)
	err = row.Scan(&book.Id, &book.Name, &book.Price)

	if err != nil {
		log.Println("Error while getting row", err)
		return book, err
	}

	log.Println("UPDATE: Name: " + s.Name + " | Price: " + s.Price)

	return book, nil
}

func (r *PostgresRepository) get(id string) (Book, error) {
	log.Println("Getting book in repo", id)

	row := r.DB.QueryRow(`SELECT * FROM demo.book WHERE id = $1`, id)
	book := Book{}
	err := row.Scan(&book.Id, &book.Name, &book.Price)

	if err != nil {
		log.Println("Error while getting row", err)
		return book, nil
	}

	return book, ErrBookNotFound

}

func (r *PostgresRepository) getAllBooks() ([]Book, error) {
	log.Println("Getting book in repo")
	//var res string
	query := `SELECT * FROM demo.book`

	book := Book{}
	var books []Book
	rows, err := r.DB.Query(query)
	if err != nil {
		fmt.Println("There was an error while querying getAllBooks")
	}
	for rows.Next() {
		err := rows.Scan(&book.Id, &book.Name, &book.Price)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func NewBookRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		DB: db,
	}
}

func (r *PostgresRepository) commitOrRollback(err error, tx *sql.Tx) error {
	if err == nil {
		if errT := tx.Commit(); errT != nil {
			return fmt.Errorf("error in tx Commit: %w", errT)
		}
	} else {
		if errT := tx.Rollback(); errT != nil {
			// choose losing rollback error type because we can use the type of the incoming error in the caller
			return fmt.Errorf("error in tx Rollback: %v : %w", errT, err)
		}
	}

	return err
}
