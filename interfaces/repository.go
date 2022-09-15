package main

import (
	"database/sql"
	"log"
	"strconv"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func save(s *Book) {
	db := dbConn()

	db.Exec("INSERT INTO Book(name, price) VALUES (s.name, s.price)")

	defer db.Close()
}

func delete(s *Book) {
	db := dbConn()
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	bookId := s.id
	del, err := db.Prepare("DELETE FROM Book WHERE id = (?)")

	if err != nil {
		panic(err.Error())
	}

	_, err = del.Exec(bookId)
	if err != nil {
		return
	}
	log.Println("DELETE")
}

func update(s *Book) {

	db := dbConn()
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	update, err := db.Prepare("UPDATE Book SET name=?, price=? WHERE id = ?")

	update.Exec(s.name, s.price)
	log.Println("UPDATE: Name: " + s.name + " | Price: " + strconv.Itoa(int(s.price)))
}
