package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_credentials = "root:bebop268170@/cursogo"
)

func main() {

	db, err := sql.Open("mysql", db_credentials)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, _ := db.Begin()

	smt, _ := tx.Prepare("insert into usuarios (id, nome) values(?,?)")

	smt.Exec(200, "Bia")
	smt.Exec(201, "Carlos")
	_, err = smt.Exec(1, "Gil")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
