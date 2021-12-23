package main

import (
	"database/sql"

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

	smt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	smt.Exec("Milton", 200)
	smt.Exec("lindomar", 3)
}
